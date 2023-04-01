import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Users } from './entity/user.entity';
import { UsersRepository } from './reposotories/UsersRepository';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(Users)
    private readonly usersRepository: UsersRepository,
  ) {}

  public findAll() {
    return this.usersRepository.find();
  }
}
