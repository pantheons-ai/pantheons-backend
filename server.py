from flask import Flask

app = Flask(__name__)


@app.route("/v1/ip", methods=["POST"])
def create_ip():
    # 1. Create folder in IPFS for this IP
    # 2. Release smart contracts on the chain
    pass


@app.route("/v1/files", methods=["POST"])
def upload_files():
    # upload files to service and then pushlish to IPFS asynchronously
    pass


@app.route("/v1/completions", methods=["POST"])
def create_completions():
    # create stories of IP by AIGC and RAG technology
    pass


