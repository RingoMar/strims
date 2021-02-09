// Code generated by Wire protocol buffer compiler, do not edit.
// Source: strims.chat.v1.OpenChatServerRequest in chat/v1/chat.proto
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

class OpenChatServerRequest(
  @field:WireField(
    tag = 1,
    adapter = "gg.strims.ppspp.proto.ChatServer#ADAPTER",
    label = WireField.Label.OMIT_IDENTITY
  )
  val server: ChatServer? = null,
  unknownFields: ByteString = ByteString.EMPTY
) : Message<OpenChatServerRequest, Nothing>(ADAPTER, unknownFields) {
  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is OpenChatServerRequest) return false
    if (unknownFields != other.unknownFields) return false
    if (server != other.server) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + server.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    if (server != null) result += """server=$server"""
    return result.joinToString(prefix = "OpenChatServerRequest{", separator = ", ", postfix = "}")
  }

  fun copy(server: ChatServer? = this.server, unknownFields: ByteString = this.unknownFields):
      OpenChatServerRequest = OpenChatServerRequest(server, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<OpenChatServerRequest> = object : ProtoAdapter<OpenChatServerRequest>(
      FieldEncoding.LENGTH_DELIMITED, 
      OpenChatServerRequest::class, 
      "type.googleapis.com/strims.chat.v1.OpenChatServerRequest", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: OpenChatServerRequest): Int {
        var size = value.unknownFields.size
        if (value.server != null) size += ChatServer.ADAPTER.encodedSizeWithTag(1, value.server)
        return size
      }

      override fun encode(writer: ProtoWriter, value: OpenChatServerRequest) {
        if (value.server != null) ChatServer.ADAPTER.encodeWithTag(writer, 1, value.server)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): OpenChatServerRequest {
        var server: ChatServer? = null
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> server = ChatServer.ADAPTER.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return OpenChatServerRequest(
          server = server,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: OpenChatServerRequest): OpenChatServerRequest = value.copy(
        server = value.server?.let(ChatServer.ADAPTER::redact),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}
