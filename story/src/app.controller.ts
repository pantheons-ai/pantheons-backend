import { Controller, Get, Post } from '@nestjs/common';
import { AppService, CreateIPRequest } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get("ping")
  getPing(): string {
    return this.appService.getPing();
  }

  @Post("ip")
  async createRootIP(request: CreateIPRequest): Promise<string> {
    const result = await this.appService.createRootIP(request)
    return result.ipId
  }
}
