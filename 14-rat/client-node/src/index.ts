import { AdminClient } from './implant_grpc_pb';
import { Command } from './implant_pb';
import * as grpc from 'grpc';


const client = new AdminClient('localhost:9090', grpc.credentials.createInsecure());
const cmd = new Command();
cmd.setIn('ls -la /d');
client.runCommand(cmd, (err, response) => {
  if (err) {
    console.error(`error:`, err);
    process.exit(1);
  }
  const out = response.getOut();
  console.log(`out`);
  console.log(out);
});