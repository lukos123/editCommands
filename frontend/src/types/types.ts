export interface CommandType {
  c: string;
  h: string;
  cf: string;
  awcf: CommandType[];
  f: CommandType[];
  na: boolean;
  o: boolean;
  n: CommandType[] | string;
}
