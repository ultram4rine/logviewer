/**
 * @fileoverview gRPC-Web generated client stub for log
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.log = require('./log_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.log.LogClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.log.LogPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.log.DHCPLogEntry,
 *   !proto.log.DHCPLogs>}
 */
const methodDescriptor_Log_GetDHCPLog = new grpc.web.MethodDescriptor(
  '/log.Log/GetDHCPLog',
  grpc.web.MethodType.UNARY,
  proto.log.DHCPLogEntry,
  proto.log.DHCPLogs,
  /**
   * @param {!proto.log.DHCPLogEntry} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.log.DHCPLogs.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.log.DHCPLogEntry,
 *   !proto.log.DHCPLogs>}
 */
const methodInfo_Log_GetDHCPLog = new grpc.web.AbstractClientBase.MethodInfo(
  proto.log.DHCPLogs,
  /**
   * @param {!proto.log.DHCPLogEntry} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.log.DHCPLogs.deserializeBinary
);


/**
 * @param {!proto.log.DHCPLogEntry} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.log.DHCPLogs)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.log.DHCPLogs>|undefined}
 *     The XHR Node Readable Stream
 */
proto.log.LogClient.prototype.getDHCPLog =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/log.Log/GetDHCPLog',
      request,
      metadata || {},
      methodDescriptor_Log_GetDHCPLog,
      callback);
};


/**
 * @param {!proto.log.DHCPLogEntry} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.log.DHCPLogs>}
 *     A native promise that resolves to the response
 */
proto.log.LogPromiseClient.prototype.getDHCPLog =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/log.Log/GetDHCPLog',
      request,
      metadata || {},
      methodDescriptor_Log_GetDHCPLog);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.log.SwName,
 *   !proto.log.Switches>}
 */
const methodDescriptor_Log_GetSimilarSwitches = new grpc.web.MethodDescriptor(
  '/log.Log/GetSimilarSwitches',
  grpc.web.MethodType.UNARY,
  proto.log.SwName,
  proto.log.Switches,
  /**
   * @param {!proto.log.SwName} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.log.Switches.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.log.SwName,
 *   !proto.log.Switches>}
 */
const methodInfo_Log_GetSimilarSwitches = new grpc.web.AbstractClientBase.MethodInfo(
  proto.log.Switches,
  /**
   * @param {!proto.log.SwName} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.log.Switches.deserializeBinary
);


/**
 * @param {!proto.log.SwName} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.log.Switches)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.log.Switches>|undefined}
 *     The XHR Node Readable Stream
 */
proto.log.LogClient.prototype.getSimilarSwitches =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/log.Log/GetSimilarSwitches',
      request,
      metadata || {},
      methodDescriptor_Log_GetSimilarSwitches,
      callback);
};


/**
 * @param {!proto.log.SwName} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.log.Switches>}
 *     A native promise that resolves to the response
 */
proto.log.LogPromiseClient.prototype.getSimilarSwitches =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/log.Log/GetSimilarSwitches',
      request,
      metadata || {},
      methodDescriptor_Log_GetSimilarSwitches);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.log.SwitchLogEntry,
 *   !proto.log.SwitchLogs>}
 */
const methodDescriptor_Log_GetSwitchLog = new grpc.web.MethodDescriptor(
  '/log.Log/GetSwitchLog',
  grpc.web.MethodType.UNARY,
  proto.log.SwitchLogEntry,
  proto.log.SwitchLogs,
  /**
   * @param {!proto.log.SwitchLogEntry} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.log.SwitchLogs.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.log.SwitchLogEntry,
 *   !proto.log.SwitchLogs>}
 */
const methodInfo_Log_GetSwitchLog = new grpc.web.AbstractClientBase.MethodInfo(
  proto.log.SwitchLogs,
  /**
   * @param {!proto.log.SwitchLogEntry} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.log.SwitchLogs.deserializeBinary
);


/**
 * @param {!proto.log.SwitchLogEntry} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.log.SwitchLogs)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.log.SwitchLogs>|undefined}
 *     The XHR Node Readable Stream
 */
proto.log.LogClient.prototype.getSwitchLog =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/log.Log/GetSwitchLog',
      request,
      metadata || {},
      methodDescriptor_Log_GetSwitchLog,
      callback);
};


/**
 * @param {!proto.log.SwitchLogEntry} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.log.SwitchLogs>}
 *     A native promise that resolves to the response
 */
proto.log.LogPromiseClient.prototype.getSwitchLog =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/log.Log/GetSwitchLog',
      request,
      metadata || {},
      methodDescriptor_Log_GetSwitchLog);
};


module.exports = proto.log;

