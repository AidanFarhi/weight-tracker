import subprocess

with open(".env.prod", "r") as env_file:
    pairs = [line.strip() for line in env_file if line.strip() and "=" in line]

all_pairs = " ".join(pairs)

subprocess.run(
    f"heroku config:set {all_pairs} --app afarhidev-weight-tracker", shell=True
)
