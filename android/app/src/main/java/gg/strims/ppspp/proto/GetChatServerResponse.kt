// Code generated by Wire protocol buffer compiler, do not edit.
// Source: strims.chat.v1.GetChatServerResponse in chat/v1/chat.proto
package gg.strims.ppspp.proto

import com.squareup.wire.FieldEncoding
import com.squareup.wire.Message
import com.squareup.wire.ProtoAdapter
import com.squareup.wire.ProtoReader
import com.squareup.wire.ProtoWriter
import com.squareup.wire.Syntax.PROTO_3
import com.squareup.wire.WireField
import kotlin.Any
import kotlin.AssertionError
import kotlin.Boolean
import kotlin.Deprecated
import kotlin.DeprecationLevel
import kotlin.Int
import kotlin.Long
import kotlin.Nothing
import kotlin.String
import kotlin.hashCode
import kotlin.jvm.JvmField
import okio.ByteString

class GetChatServerResponse(
  @field:WireField(
    tag = 1,
    adapter = "gg.strims.ppspp.proto.ChatServer#ADAPTER",
    label = WireField.Label.OMIT_IDENTITY,
    jsonName = "chatServer"
  )
  val chat_server: ChatServer? = null,
  unknownFields: ByteString = ByteString.EMPTY
) : Message<GetChatServerResponse, Nothing>(ADAPTER, unknownFields) {
  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is GetChatServerResponse) return false
    if (unknownFields != other.unknownFields) return false
    if (chat_server != other.chat_server) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + chat_server.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    if (chat_server != null) result += """chat_server=$chat_server"""
    return result.joinToString(prefix = "GetChatServerResponse{", separator = ", ", postfix = "}")
  }

  fun copy(chat_server: ChatServer? = this.chat_server, unknownFields: ByteString =
      this.unknownFields): GetChatServerResponse = GetChatServerResponse(chat_server, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<GetChatServerResponse> = object : ProtoAdapter<GetChatServerResponse>(
      FieldEncoding.LENGTH_DELIMITED, 
      GetChatServerResponse::class, 
      "type.googleapis.com/strims.chat.v1.GetChatServerResponse", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: GetChatServerResponse): Int {
        var size = value.unknownFields.size
        if (value.chat_server != null) size += ChatServer.ADAPTER.encodedSizeWithTag(1,
            value.chat_server)
        return size
      }

      override fun encode(writer: ProtoWriter, value: GetChatServerResponse) {
        if (value.chat_server != null) ChatServer.ADAPTER.encodeWithTag(writer, 1,
            value.chat_server)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): GetChatServerResponse {
        var chat_server: ChatServer? = null
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> chat_server = ChatServer.ADAPTER.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return GetChatServerResponse(
          chat_server = chat_server,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: GetChatServerResponse): GetChatServerResponse = value.copy(
        chat_server = value.chat_server?.let(ChatServer.ADAPTER::redact),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}
