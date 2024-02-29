from langchain_community.document_loaders import PyPDFLoader

loader = PyPDFLoader("./test/Ethereum_Whitepaper_Buterin_2014.pdf")

pages = loader.load_and_split()

