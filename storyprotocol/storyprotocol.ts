import { privateKeyToAccount } from "viem/accounts";
import { StoryClient, StoryConfig } from "@story-protocol/core-sdk";
import { http } from 'viem';


const WALLET_PRIVATE_KEY = process.env.WALLET_PRIVATE_KEY || "0x";
const account = privateKeyToAccount(WALLET_PRIVATE_KEY as `0x${string}`);



const config: StoryConfig = {
        transport: http(process.env.RPC_PROVIDER_URL),
        account: account
};
const client = StoryClient.newClient(config);