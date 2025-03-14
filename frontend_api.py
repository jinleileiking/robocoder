import Agently
import typer
import json
from loguru import logger
from utils import workflow_start
from utils import camel_to_snake
from agent import Agent
from agent import write_file


agent = Agent(debug=True)

workflow = Agently.Workflow()
# workflow.settings.set("is_debug", True)
ApiDocImplementWorkflow = Agently.Workflow()


@ApiDocImplementWorkflow.chunk()
def api_doc(inputs, storage):

    project = Agently.Workflow.public_storage.get("prj_name")

    struct = camel_to_snake(inputs["default"])
    with open(f"{project}/backend/struct_{struct}.go", "r") as file:
        struct_def = file.read()

    with open(f"{project}/backend/router_{struct}.go", "r") as file:
        router_implement = file.read()

    agent.input(
        {
            "struct_def": struct_def,
            "router_implement": router_implement,
        }
    )

    reply = agent.run_prompt(
        "prompt/frontend/api_doc.prompt",
    )

    write_file(f"api_{struct}.md", reply, dir="frontend/docs")
    storage.set("struct", struct)

    return reply


# @ApiDocImplementWorkflow.chunk()
# def interface_definition(inputs, storage):
#     struct = storage.get("struct")
#     logger.info(f"storage: {storage}")
#     agent.input(
#         {
#             "api_doc": inputs["default"],
#         }
#     )
#     reply = agent.run_prompt("prompt/frontend/interface_definition.prompt")
#     write_file(f"{struct}.ts", reply, dir="frontend/src/api")
#
#     return reply


@ApiDocImplementWorkflow.chunk()
def interface_implement(inputs, storage):
    struct = storage.get("struct")
    logger.info(f"storage: {storage}")
    logger.info(f"inputs['default']: {inputs['default']}")

    agent.input(
        {
            "doc": inputs["default"],
        }
    )
    reply = agent.run_prompt("prompt/frontend/interface_implement.prompt")
    write_file(f"{struct}.ts", reply, dir="frontend/src/api")

    return reply


(workflow.loop_with(ApiDocImplementWorkflow).connect_to("END"))

(
    ApiDocImplementWorkflow.connect_to("api_doc")
    # .connect_to("interface_definition")
    .connect_to("interface_implement").connect_to("END")
)


def main(project: str):

    with open(f"{project}/structlist.json", "r") as file:
        structlist = file.read()

    Agently.Workflow.public_storage.set("structlist", structlist)
    Agently.Workflow.public_storage.set("prj_name", project)

    write_file("frontend_api_workflow.mermaid", workflow.draw())

    workflow_start(workflow, start=json.loads(structlist))


if __name__ == "__main__":
    typer.run(main)
