import Agently
import httpx
import re
from string import Template
from enum import Enum, auto
from pathlib import Path

from loguru import logger

from utils import run_agent


def write_file(filename: str, content, dir=None):
    prj_name = Agently.Workflow.public_storage.get("prj_name")
    if dir is not None:
        directory_path = Path(f"{prj_name}/{dir}")
        directory_path.mkdir(parents=True, exist_ok=True)
        file_path = f"{prj_name}/{dir}/{filename}"
    else:
        file_path = f"{prj_name}/{filename}"

    try:
        with open(file_path, "w") as file:
            file.write(str(content))
        logger.info(f"Content written to {filename} successfully.")
    except IOError as e:
        logger.error(f"An error occurred: {e}")


class Agent:
    class PromptPosition(Enum):
        START = auto()
        END = auto()

    def __init__(self, debug=None):
        transport = httpx.AsyncHTTPTransport(retries=10)

        client = httpx.AsyncClient(timeout=None, transport=transport)

        agent_factory = Agently.AgentFactory(is_debug=False)

        agent_factory.set_settings("current_model", "AzureOpenAI").set_settings(
            "model.AzureOpenAI.auth",
            {
                "api_key": "yourkey",
                "api_version": "version",
                "azure_endpoint": "https://xxxxxx.openai.azure.com",
                "azure_deployment": "gtp4o-0801",
            },
        ).set_settings(
            "model.AzureOpenAI.options", {"model": "gpt-4o", "temperature": 0.4}
        ).set_settings(
            "model.AzureOpenAI.httpx_client",
            client,
        )

        agent = agent_factory.create_agent()

        if debug == True:
            agent.set_settings("is_debug", True)
        self.agent = agent

    def input(self, input):
        self.agent = self.agent.input(input)

    def run_prompt(
        self, prompt_file: str, position: PromptPosition = None, output=None
    ):
        with open(prompt_file, "r") as file:
            prompt_template = file.read()

        prompt = render(prompt_template, position)
        reply = run_agent(self.agent, prompt, output)

        match = re.search(r"/([^/]+)\.prompt$", prompt_file)
        if match:
            prompt_name = match.group(1)
        else:
            logger.error(f"{prompt_name}, No match found")

        logger.info(f"setting storage {prompt_name}")
        Agently.Workflow.public_storage.set(prompt_name, reply)
        return reply


def render(content: str, position: Agent.PromptPosition = None):
    template = Template(content)
    result = template.safe_substitute(
        prj_name=Agently.Workflow.public_storage.get("prj_name"),
        # backend
        desc=Agently.Workflow.public_storage.get("desc"),
        prd_content=Agently.Workflow.public_storage.get("prd"),
        functionlist=Agently.Workflow.public_storage.get("funclist"),
        struct=Agently.Workflow.public_storage.get("struct"),
        structlist=Agently.Workflow.public_storage.get("structlist"),
        # routerlist=Agently.Workflow.public_storage.get("routerlist"),
        # frontend
        origin_demond=Agently.Workflow.public_storage.get("origin_demond"),
        modules=Agently.Workflow.public_storage.get("ui_design"),
        page=Agently.Workflow.public_storage.get("page"),
    )

    if position == Agent.PromptPosition.START:
        result = "<|im_start|>\n" + result
    if position == Agent.PromptPosition.END:
        result = result + "<|im_end|>"

    # print("--------------------PROMPT-------------------------")
    # print(result)
    # print("---------------------------------------------------")
    return result
