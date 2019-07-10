/**
 * @fileoverview gRPC-Web generated client stub for chat
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.chat = require('./chat_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.chat.ChatClient =
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
proto.chat.ChatPromiseClient =
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
 *   !proto.chat.ConnReq,
 *   !proto.chat.Msg>}
 */
const methodDescriptor_Chat_Conn = new grpc.web.MethodDescriptor(
  '/chat.Chat/Conn',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.chat.ConnReq,
  proto.chat.Msg,
  /** @param {!proto.chat.ConnReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.Msg.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.chat.ConnReq,
 *   !proto.chat.Msg>}
 */
const methodInfo_Chat_Conn = new grpc.web.AbstractClientBase.MethodInfo(
  proto.chat.Msg,
  /** @param {!proto.chat.ConnReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.Msg.deserializeBinary
);


/**
 * @param {!proto.chat.ConnReq} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.chat.Msg>}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatClient.prototype.conn =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/chat.Chat/Conn',
      request,
      metadata || {},
      methodDescriptor_Chat_Conn);
};


/**
 * @param {!proto.chat.ConnReq} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.chat.Msg>}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatPromiseClient.prototype.conn =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/chat.Chat/Conn',
      request,
      metadata || {},
      methodDescriptor_Chat_Conn);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.chat.PubMsgReq,
 *   !proto.chat.PubMsgRsp>}
 */
const methodDescriptor_Chat_PubMsg = new grpc.web.MethodDescriptor(
  '/chat.Chat/PubMsg',
  grpc.web.MethodType.UNARY,
  proto.chat.PubMsgReq,
  proto.chat.PubMsgRsp,
  /** @param {!proto.chat.PubMsgReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.PubMsgRsp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.chat.PubMsgReq,
 *   !proto.chat.PubMsgRsp>}
 */
const methodInfo_Chat_PubMsg = new grpc.web.AbstractClientBase.MethodInfo(
  proto.chat.PubMsgRsp,
  /** @param {!proto.chat.PubMsgReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.PubMsgRsp.deserializeBinary
);


/**
 * @param {!proto.chat.PubMsgReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.chat.PubMsgRsp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.chat.PubMsgRsp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatClient.prototype.pubMsg =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/chat.Chat/PubMsg',
      request,
      metadata || {},
      methodDescriptor_Chat_PubMsg,
      callback);
};


/**
 * @param {!proto.chat.PubMsgReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.chat.PubMsgRsp>}
 *     A native promise that resolves to the response
 */
proto.chat.ChatPromiseClient.prototype.pubMsg =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/chat.Chat/PubMsg',
      request,
      metadata || {},
      methodDescriptor_Chat_PubMsg);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.chat.HeartBeatReq,
 *   !proto.chat.HeartBeatRsp>}
 */
const methodDescriptor_Chat_HeartBeat = new grpc.web.MethodDescriptor(
  '/chat.Chat/HeartBeat',
  grpc.web.MethodType.UNARY,
  proto.chat.HeartBeatReq,
  proto.chat.HeartBeatRsp,
  /** @param {!proto.chat.HeartBeatReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.HeartBeatRsp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.chat.HeartBeatReq,
 *   !proto.chat.HeartBeatRsp>}
 */
const methodInfo_Chat_HeartBeat = new grpc.web.AbstractClientBase.MethodInfo(
  proto.chat.HeartBeatRsp,
  /** @param {!proto.chat.HeartBeatReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.HeartBeatRsp.deserializeBinary
);


/**
 * @param {!proto.chat.HeartBeatReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.chat.HeartBeatRsp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.chat.HeartBeatRsp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatClient.prototype.heartBeat =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/chat.Chat/HeartBeat',
      request,
      metadata || {},
      methodDescriptor_Chat_HeartBeat,
      callback);
};


/**
 * @param {!proto.chat.HeartBeatReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.chat.HeartBeatRsp>}
 *     A native promise that resolves to the response
 */
proto.chat.ChatPromiseClient.prototype.heartBeat =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/chat.Chat/HeartBeat',
      request,
      metadata || {},
      methodDescriptor_Chat_HeartBeat);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.chat.GroupActionReq,
 *   !proto.chat.GroupActionRsp>}
 */
const methodDescriptor_Chat_GroupAction = new grpc.web.MethodDescriptor(
  '/chat.Chat/GroupAction',
  grpc.web.MethodType.UNARY,
  proto.chat.GroupActionReq,
  proto.chat.GroupActionRsp,
  /** @param {!proto.chat.GroupActionReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.GroupActionRsp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.chat.GroupActionReq,
 *   !proto.chat.GroupActionRsp>}
 */
const methodInfo_Chat_GroupAction = new grpc.web.AbstractClientBase.MethodInfo(
  proto.chat.GroupActionRsp,
  /** @param {!proto.chat.GroupActionReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.GroupActionRsp.deserializeBinary
);


/**
 * @param {!proto.chat.GroupActionReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.chat.GroupActionRsp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.chat.GroupActionRsp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatClient.prototype.groupAction =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/chat.Chat/GroupAction',
      request,
      metadata || {},
      methodDescriptor_Chat_GroupAction,
      callback);
};


/**
 * @param {!proto.chat.GroupActionReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.chat.GroupActionRsp>}
 *     A native promise that resolves to the response
 */
proto.chat.ChatPromiseClient.prototype.groupAction =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/chat.Chat/GroupAction',
      request,
      metadata || {},
      methodDescriptor_Chat_GroupAction);
};


module.exports = proto.chat;

