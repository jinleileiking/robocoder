import Agently
import typer
import json
import os
import shutil
import subprocess
from pathlib import Path

from loguru import logger

from agent import write_file
from agent import Agent
from utils import camel_to_snake
from utils import workflow_start

agent = Agent()
workflow = Agently.Workflow()

StructWorkflow = Agently.Workflow()
RouterWorkflow = Agently.Workflow()


@workflow.chunk()
def prd(inputs, storage):
    reply = agent.run_prompt("prompt/backend/prd.prompt", Agent.PromptPosition.START)

    write_file("prd.md", reply)
    return reply


@workflow.chunk()
def funclist(inputs, storage):
    reply = agent.run_prompt("prompt/backend/funclist.prompt")
    return reply


@workflow.chunk()
def structlist(inputs, storage):
    reply = agent.run_prompt("prompt/backend/structlist.prompt")

    write_file("structlist.json", reply)

    data = json.loads(reply)

    if isinstance(data, list):
        return data
    else:
        raise RuntimeError(f"structlist not return list, reply: {reply}")

    return reply


@workflow.chunk()
def main_implement(inputs, storage):
    reply = agent.run_prompt("prompt/backend/main.prompt", Agent.PromptPosition.END)

    write_file("main.go", reply, dir="backend")
    return reply


@StructWorkflow.chunk()
def struct_def(inputs, storage):

    struct_name = inputs["default"]
    logger.info(f"generating struct {struct_name}")
    Agently.Workflow.public_storage.set("struct", struct_name)

    reply = agent.run_prompt(
        "prompt/backend/struct_def.prompt",
    )

    write_file(f"struct_{camel_to_snake(struct_name)}.go", reply, dir="backend")
    return reply


@RouterWorkflow.chunk()
def router_implement(inputs, storage):

    struct_name = inputs["default"]
    logger.info(f"generating router {struct_name}")
    Agently.Workflow.public_storage.set("struct", struct_name)

    reply = agent.run_prompt(
        "prompt/backend/router_imp.prompt",
    )

    write_file(f"router_{camel_to_snake(struct_name)}.go", reply, dir="backend")
    return reply


(StructWorkflow.connect_to("struct_def").connect_to("END"))


(RouterWorkflow.connect_to("router_implement").connect_to("END"))

(
    workflow.connect_to("prd")
    .connect_to("funclist")
    .connect_to("structlist")
    .loop_with(StructWorkflow)
    .connect_to("END")
)


workflow.chunks["structlist"].loop_with(RouterWorkflow).connect_to("END")
workflow.chunks["structlist"].connect_to(main_implement).connect_to("END")


def main(project: str):

    directory_path = Path(f"./{project}/backend")
    directory_path.mkdir(parents=True, exist_ok=True)

    shutil.copy("./go_examples/db.go", f"./{project}/backend/db.go")

    # replace_package_name(f"./{project}/db.go", "package main", f"package {project}")

    os.chdir(f"./{project}/backend")
    subprocess.run(["go", "mod", "init", project], capture_output=True, text=True)
    os.chdir("../..")

    with open("prd.txt", "r") as file:
        prd = file.read()

    Agently.Workflow.public_storage.set("desc", prd)
    Agently.Workflow.public_storage.set("prj_name", project)

    write_file("backend_workflow.mermaid", workflow.draw())

    workflow_start(workflow)


if __name__ == "__main__":
    typer.run(main)
