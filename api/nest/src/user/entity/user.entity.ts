import { Column, Entity, ObjectID, ObjectIdColumn } from 'typeorm';

@Entity('users')
export class Users {
  @ObjectIdColumn()
  id: ObjectID;
  @Column()
  name: string;
  @Column()
  gender: string;
  @Column()
  age: number;
}
