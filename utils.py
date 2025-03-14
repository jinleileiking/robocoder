import re
import time
from loguru import logger


def workflow_start(workflow, start=None):
    start_time = time.time()

    workflow.start(start)

    end_time = time.time()
    execution_time = end_time - start_time
    logger.info(f"cost: {execution_time:.4f}s")


# def camel_to_snake(name):
#     # 在每个大写字母前面加上下划线，然后转换为小写
#     snake_case = re.sub(
#         r"(?<!^)(?<!_)(?<![A-Z])(?<!\d)(?<!\s)(?<!\d\s)(?<![\d\s])[A-Z]",
#         r"_\g<0>",
#         name,
#     ).lower()
#     return snake_case


def camel_to_snake(name):
    # 使用正则表达式找到大写字母，并在这些字母前加上下划线
    s1 = re.sub("(.)([A-Z][a-z]+)", r"\1_\2", name)
    # 将首字母大写转换为小写
    return re.sub("([a-z0-9])([A-Z])", r"\1_\2", s1).lower()


def run_agent(agent, prompt, output=None, start=None):
    while True:
        if output is not None:
            instance = agent.instruct(prompt).output(output)
        else:
            instance = agent.instruct(prompt)

        if start is not None:
            reply = instance.start(start)
        else:
            reply = instance.start()

        if reply is not None:
            return reply


# def replace_package_name(file_path, old_package_name, new_package_name):
#     with open(file_path, "r") as file:
#         file_content = file.read()
#
#     updated_content = re.sub(rf"\b{old_package_name}\b", new_package_name, file_content)
#
#     with open(file_path, "w") as file:
#         file.write(updated_content)
