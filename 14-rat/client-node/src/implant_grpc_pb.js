// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// implant.proto
'use strict';
var grpc = require('grpc');
var implant_pb = require('./implant_pb.js');

function serialize_grpcapi_Command(arg) {
  if (!(arg instanceof implant_pb.Command)) {
    throw new Error('Expected argument of type grpcapi.Command');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_grpcapi_Command(buffer_arg) {
  return implant_pb.Command.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_grpcapi_Empty(arg) {
  if (!(arg instanceof implant_pb.Empty)) {
    throw new Error('Expected argument of type grpcapi.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_grpcapi_Empty(buffer_arg) {
  return implant_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}


// Implant defines our C2 API functions
var ImplantService = exports.ImplantService = {
  fetchCommand: {
    path: '/grpcapi.Implant/FetchCommand',
    requestStream: false,
    responseStream: false,
    requestType: implant_pb.Empty,
    responseType: implant_pb.Command,
    requestSerialize: serialize_grpcapi_Empty,
    requestDeserialize: deserialize_grpcapi_Empty,
    responseSerialize: serialize_grpcapi_Command,
    responseDeserialize: deserialize_grpcapi_Command,
  },
  sendOutput: {
    path: '/grpcapi.Implant/SendOutput',
    requestStream: false,
    responseStream: false,
    requestType: implant_pb.Command,
    responseType: implant_pb.Empty,
    requestSerialize: serialize_grpcapi_Command,
    requestDeserialize: deserialize_grpcapi_Command,
    responseSerialize: serialize_grpcapi_Empty,
    responseDeserialize: deserialize_grpcapi_Empty,
  },
};

exports.ImplantClient = grpc.makeGenericClientConstructor(ImplantService);
// Admin defines our Admin API functions
var AdminService = exports.AdminService = {
  runCommand: {
    path: '/grpcapi.Admin/RunCommand',
    requestStream: false,
    responseStream: false,
    requestType: implant_pb.Command,
    responseType: implant_pb.Command,
    requestSerialize: serialize_grpcapi_Command,
    requestDeserialize: deserialize_grpcapi_Command,
    responseSerialize: serialize_grpcapi_Command,
    responseDeserialize: deserialize_grpcapi_Command,
  },
};

exports.AdminClient = grpc.makeGenericClientConstructor(AdminService);
