import { Injectable } from '@nestjs/common';
import { StoryClient, StoryConfig } from '@story-protocol/core-sdk'
import { http, Address } from 'viem'
import { privateKeyToAccount } from 'viem/accounts'

export type CreateIPRequest = {
  tokenContractAddress: `0x${string}`;
  tokenId: string;
  policyId?: string;
  ipName?: string;
  contentHash?: `0x${string}`;
  uri?: string;
};

export type CreateIPResponse = {
  txHash?: string;
  ipId?: string;
};

@Injectable()
export class AppService {
  getPing(): string {
    return 'pong';
  }

  async createRootIP(request: CreateIPRequest): Promise<CreateIPResponse> {
    const WALLET_PRIVATE_KEY = process.env.WALLET_PRIVATE_KEY || "0x";
    const account = privateKeyToAccount(WALLET_PRIVATE_KEY as `0x${string}`);
    const config: StoryConfig = {
        account,
        transport: http(process.env.RPC_PROVIDER_URL),
    }
    const client = StoryClient.newClient(config)

    const registeredIpAsset = await client.ipAsset.registerRootIp({
        tokenContractAddress: request.tokenContractAddress,
        tokenId: request.tokenId,
        policyId: '0',
        txOptions: { waitForTransaction: true }
    })
    console.log(`Root IPA created at transaction hash ${registeredIpAsset.txHash}, IPA ID: ${registeredIpAsset.ipId}`)

    return registeredIpAsset
  }
}
