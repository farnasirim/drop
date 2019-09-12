/**
 * @fileoverview gRPC-Web generated client stub for drop
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.drop = require('./drop_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.drop.DropApiClient =
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

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.drop.DropApiPromiseClient =
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

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.drop.LoginRequest,
 *   !proto.drop.LoginResponse>}
 */
const methodDescriptor_DropApi_TwoStepLogin = new grpc.web.MethodDescriptor(
  '/drop.DropApi/TwoStepLogin',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.drop.LoginRequest,
  proto.drop.LoginResponse,
  /** @param {!proto.drop.LoginRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.drop.LoginResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.drop.LoginRequest,
 *   !proto.drop.LoginResponse>}
 */
const methodInfo_DropApi_TwoStepLogin = new grpc.web.AbstractClientBase.MethodInfo(
  proto.drop.LoginResponse,
  /** @param {!proto.drop.LoginRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.drop.LoginResponse.deserializeBinary
);


/**
 * @param {!proto.drop.LoginRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.drop.LoginResponse>}
 *     The XHR Node Readable Stream
 */
proto.drop.DropApiClient.prototype.twoStepLogin =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/drop.DropApi/TwoStepLogin',
      request,
      metadata || {},
      methodDescriptor_DropApi_TwoStepLogin);
};


/**
 * @param {!proto.drop.LoginRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.drop.LoginResponse>}
 *     The XHR Node Readable Stream
 */
proto.drop.DropApiPromiseClient.prototype.twoStepLogin =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/drop.DropApi/TwoStepLogin',
      request,
      metadata || {},
      methodDescriptor_DropApi_TwoStepLogin);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.drop.PutLinkRequest,
 *   !proto.drop.PutLinkResponse>}
 */
const methodDescriptor_DropApi_PutLink = new grpc.web.MethodDescriptor(
  '/drop.DropApi/PutLink',
  grpc.web.MethodType.UNARY,
  proto.drop.PutLinkRequest,
  proto.drop.PutLinkResponse,
  /** @param {!proto.drop.PutLinkRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.drop.PutLinkResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.drop.PutLinkRequest,
 *   !proto.drop.PutLinkResponse>}
 */
const methodInfo_DropApi_PutLink = new grpc.web.AbstractClientBase.MethodInfo(
  proto.drop.PutLinkResponse,
  /** @param {!proto.drop.PutLinkRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.drop.PutLinkResponse.deserializeBinary
);


/**
 * @param {!proto.drop.PutLinkRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.drop.PutLinkResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.drop.PutLinkResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.drop.DropApiClient.prototype.putLink =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/drop.DropApi/PutLink',
      request,
      metadata || {},
      methodDescriptor_DropApi_PutLink,
      callback);
};


/**
 * @param {!proto.drop.PutLinkRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.drop.PutLinkResponse>}
 *     A native promise that resolves to the response
 */
proto.drop.DropApiPromiseClient.prototype.putLink =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/drop.DropApi/PutLink',
      request,
      metadata || {},
      methodDescriptor_DropApi_PutLink);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.drop.RemoveLinkRequest,
 *   !proto.drop.RemoveLinkResponse>}
 */
const methodDescriptor_DropApi_RemoveLink = new grpc.web.MethodDescriptor(
  '/drop.DropApi/RemoveLink',
  grpc.web.MethodType.UNARY,
  proto.drop.RemoveLinkRequest,
  proto.drop.RemoveLinkResponse,
  /** @param {!proto.drop.RemoveLinkRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.drop.RemoveLinkResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.drop.RemoveLinkRequest,
 *   !proto.drop.RemoveLinkResponse>}
 */
const methodInfo_DropApi_RemoveLink = new grpc.web.AbstractClientBase.MethodInfo(
  proto.drop.RemoveLinkResponse,
  /** @param {!proto.drop.RemoveLinkRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.drop.RemoveLinkResponse.deserializeBinary
);


/**
 * @param {!proto.drop.RemoveLinkRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.drop.RemoveLinkResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.drop.RemoveLinkResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.drop.DropApiClient.prototype.removeLink =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/drop.DropApi/RemoveLink',
      request,
      metadata || {},
      methodDescriptor_DropApi_RemoveLink,
      callback);
};


/**
 * @param {!proto.drop.RemoveLinkRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.drop.RemoveLinkResponse>}
 *     A native promise that resolves to the response
 */
proto.drop.DropApiPromiseClient.prototype.removeLink =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/drop.DropApi/RemoveLink',
      request,
      metadata || {},
      methodDescriptor_DropApi_RemoveLink);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.drop.GetLinksRequest,
 *   !proto.drop.GetLinksResponse>}
 */
const methodDescriptor_DropApi_GetLinks = new grpc.web.MethodDescriptor(
  '/drop.DropApi/GetLinks',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.drop.GetLinksRequest,
  proto.drop.GetLinksResponse,
  /** @param {!proto.drop.GetLinksRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.drop.GetLinksResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.drop.GetLinksRequest,
 *   !proto.drop.GetLinksResponse>}
 */
const methodInfo_DropApi_GetLinks = new grpc.web.AbstractClientBase.MethodInfo(
  proto.drop.GetLinksResponse,
  /** @param {!proto.drop.GetLinksRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.drop.GetLinksResponse.deserializeBinary
);


/**
 * @param {!proto.drop.GetLinksRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.drop.GetLinksResponse>}
 *     The XHR Node Readable Stream
 */
proto.drop.DropApiClient.prototype.getLinks =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/drop.DropApi/GetLinks',
      request,
      metadata || {},
      methodDescriptor_DropApi_GetLinks);
};


/**
 * @param {!proto.drop.GetLinksRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.drop.GetLinksResponse>}
 *     The XHR Node Readable Stream
 */
proto.drop.DropApiPromiseClient.prototype.getLinks =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/drop.DropApi/GetLinks',
      request,
      metadata || {},
      methodDescriptor_DropApi_GetLinks);
};


module.exports = proto.drop;

