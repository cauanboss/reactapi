import { Repository } from 'typeorm';
import { Users } from '../entity/user.entity';

export class UsersRepository extends Repository<Users> {}
