import Agently
import typer
from loguru import logger
from utils import workflow_start
from agent import Agent
from agent import write_file


from pathlib import Path

agent = Agent()

workflow = Agently.Workflow()

PageImplementWorkflow = Agently.Workflow()
RouterImplementWorkflow = Agently.Workflow()


@workflow.chunk()
def ui_design(inputs, storage):
    reply = agent.run_prompt(
        "prompt/frontend/ui_design.prompt", Agent.PromptPosition.START
    )
    write_file("ui_design.json", reply, dir="frontend")
    # Agently.Workflow.public_storage.set("ui_design", reply)

    return reply


@workflow.chunk()
def page(inputs, storage):
    reply = agent.run_prompt(
        "prompt/frontend/page.prompt",
        output={
            "pages": (
                [
                    ("str", "返回的page"),
                    ("str", "返回的content"),
                ],
            ),
        },
    )
    write_file("page.json", reply, dir="frontend")
    return reply["pages"]


@PageImplementWorkflow.chunk()
def page_implement(inputs, storage):
    input = inputs["default"]
    logger.info(f"page_implement input: {input}")
    Agently.Workflow.public_storage.set("page", input)

    reply = agent.run_prompt("prompt/frontend/page_implement.prompt")
    page = input["page"]
    write_file(f"{page}.vue", reply, dir="frontend/src/pages/")

    return reply


# @RouterImplementWorkflow.chunk()
@workflow.chunk()
def router_implement(inputs, storage):
    input = inputs["default"]
    logger.info(f"router_implement input: {input}")
    Agently.Workflow.public_storage.set("page", input)

    reply = agent.run_prompt("prompt/frontend/router_implement.prompt")
    write_file("index.ts", reply, dir="frontend/src/router/")

    return reply


(workflow.connect_to("ui_design").connect_to("page"))

# (workflow.chunks["page"].loop_with(RouterImplementWorkflow).connect_to("END"))
(workflow.chunks["page"].loop_with(PageImplementWorkflow).connect_to("END"))
(workflow.chunks["page"].connect_to("router_implement").connect_to("END"))

(PageImplementWorkflow.connect_to("page_implement").connect_to("END"))
# (RouterImplementWorkflow.connect_to("router_implement").connect_to("END"))


def main(project: str):

    directory_path = Path(f"{project}/frontend")
    directory_path.mkdir(parents=True, exist_ok=True)

    with open("prd.txt", "r") as file:
        prd = file.read()

    Agently.Workflow.public_storage.set("origin_demond", prd)
    Agently.Workflow.public_storage.set("prj_name", project)

    write_file("frontend_ui_workflow.mermaid", workflow.draw())

    workflow_start(workflow)


if __name__ == "__main__":
    typer.run(main)
