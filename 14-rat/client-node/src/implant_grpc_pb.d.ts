// package: grpcapi
// file: implant.proto

import * as grpc from 'grpc';
import * as implant_pb from './implant_pb';

interface IImplantService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  fetchCommand: IImplantService_IFetchCommand;
  sendOutput: IImplantService_ISendOutput;
}

interface IImplantService_IFetchCommand {
  path: string; // "/grpcapi.Implant/FetchCommand"
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<implant_pb.Empty>;
  requestDeserialize: grpc.deserialize<implant_pb.Empty>;
  responseSerialize: grpc.serialize<implant_pb.Command>;
  responseDeserialize: grpc.deserialize<implant_pb.Command>;
}

interface IImplantService_ISendOutput {
  path: string; // "/grpcapi.Implant/SendOutput"
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<implant_pb.Command>;
  requestDeserialize: grpc.deserialize<implant_pb.Command>;
  responseSerialize: grpc.serialize<implant_pb.Empty>;
  responseDeserialize: grpc.deserialize<implant_pb.Empty>;
}

export const ImplantService: IImplantService;
export interface IImplantServer {
  fetchCommand: grpc.handleUnaryCall<implant_pb.Empty, implant_pb.Command>;
  sendOutput: grpc.handleUnaryCall<implant_pb.Command, implant_pb.Empty>;
}

export interface IImplantClient {
  fetchCommand(request: implant_pb.Empty, callback: (error: Error | null, response: implant_pb.Command) => void): grpc.ClientUnaryCall;
  fetchCommand(request: implant_pb.Empty, metadata: grpc.Metadata, callback: (error: Error | null, response: implant_pb.Command) => void): grpc.ClientUnaryCall;
  sendOutput(request: implant_pb.Command, callback: (error: Error | null, response: implant_pb.Empty) => void): grpc.ClientUnaryCall;
  sendOutput(request: implant_pb.Command, metadata: grpc.Metadata, callback: (error: Error | null, response: implant_pb.Empty) => void): grpc.ClientUnaryCall;
}

export class ImplantClient extends grpc.Client implements IImplantClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  public fetchCommand(request: implant_pb.Empty, callback: (error: Error | null, response: implant_pb.Command) => void): grpc.ClientUnaryCall;
  public fetchCommand(request: implant_pb.Empty, metadata: grpc.Metadata, callback: (error: Error | null, response: implant_pb.Command) => void): grpc.ClientUnaryCall;
  public sendOutput(request: implant_pb.Command, callback: (error: Error | null, response: implant_pb.Empty) => void): grpc.ClientUnaryCall;
  public sendOutput(request: implant_pb.Command, metadata: grpc.Metadata, callback: (error: Error | null, response: implant_pb.Empty) => void): grpc.ClientUnaryCall;
}

interface IAdminService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  runCommand: IAdminService_IRunCommand;
}

interface IAdminService_IRunCommand {
  path: string; // "/grpcapi.Admin/RunCommand"
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<implant_pb.Command>;
  requestDeserialize: grpc.deserialize<implant_pb.Command>;
  responseSerialize: grpc.serialize<implant_pb.Command>;
  responseDeserialize: grpc.deserialize<implant_pb.Command>;
}

export const AdminService: IAdminService;
export interface IAdminServer {
  runCommand: grpc.handleUnaryCall<implant_pb.Command, implant_pb.Command>;
}

export interface IAdminClient {
  runCommand(request: implant_pb.Command, callback: (error: Error | null, response: implant_pb.Command) => void): grpc.ClientUnaryCall;
  runCommand(request: implant_pb.Command, metadata: grpc.Metadata, callback: (error: Error | null, response: implant_pb.Command) => void): grpc.ClientUnaryCall;
}

export class AdminClient extends grpc.Client implements IAdminClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  public runCommand(request: implant_pb.Command, callback: (error: Error | null, response: implant_pb.Command) => void): grpc.ClientUnaryCall;
  public runCommand(request: implant_pb.Command, metadata: grpc.Metadata, callback: (error: Error | null, response: implant_pb.Command) => void): grpc.ClientUnaryCall;
}

