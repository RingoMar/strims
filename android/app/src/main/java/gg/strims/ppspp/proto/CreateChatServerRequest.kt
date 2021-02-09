// Code generated by Wire protocol buffer compiler, do not edit.
// Source: strims.chat.v1.CreateChatServerRequest in chat/v1/chat.proto
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

class CreateChatServerRequest(
  @field:WireField(
    tag = 2,
    adapter = "com.squareup.wire.ProtoAdapter#BYTES",
    label = WireField.Label.OMIT_IDENTITY,
    jsonName = "networkKey"
  )
  val network_key: ByteString = ByteString.EMPTY,
  @field:WireField(
    tag = 3,
    adapter = "gg.strims.ppspp.proto.ChatRoom#ADAPTER",
    label = WireField.Label.OMIT_IDENTITY,
    jsonName = "chatRoom"
  )
  val chat_room: ChatRoom? = null,
  unknownFields: ByteString = ByteString.EMPTY
) : Message<CreateChatServerRequest, Nothing>(ADAPTER, unknownFields) {
  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is CreateChatServerRequest) return false
    if (unknownFields != other.unknownFields) return false
    if (network_key != other.network_key) return false
    if (chat_room != other.chat_room) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + network_key.hashCode()
      result = result * 37 + chat_room.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    result += """network_key=$network_key"""
    if (chat_room != null) result += """chat_room=$chat_room"""
    return result.joinToString(prefix = "CreateChatServerRequest{", separator = ", ", postfix = "}")
  }

  fun copy(
    network_key: ByteString = this.network_key,
    chat_room: ChatRoom? = this.chat_room,
    unknownFields: ByteString = this.unknownFields
  ): CreateChatServerRequest = CreateChatServerRequest(network_key, chat_room, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<CreateChatServerRequest> = object :
        ProtoAdapter<CreateChatServerRequest>(
      FieldEncoding.LENGTH_DELIMITED, 
      CreateChatServerRequest::class, 
      "type.googleapis.com/strims.chat.v1.CreateChatServerRequest", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: CreateChatServerRequest): Int {
        var size = value.unknownFields.size
        if (value.network_key != ByteString.EMPTY) size += ProtoAdapter.BYTES.encodedSizeWithTag(2,
            value.network_key)
        if (value.chat_room != null) size += ChatRoom.ADAPTER.encodedSizeWithTag(3, value.chat_room)
        return size
      }

      override fun encode(writer: ProtoWriter, value: CreateChatServerRequest) {
        if (value.network_key != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 2,
            value.network_key)
        if (value.chat_room != null) ChatRoom.ADAPTER.encodeWithTag(writer, 3, value.chat_room)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): CreateChatServerRequest {
        var network_key: ByteString = ByteString.EMPTY
        var chat_room: ChatRoom? = null
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            2 -> network_key = ProtoAdapter.BYTES.decode(reader)
            3 -> chat_room = ChatRoom.ADAPTER.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return CreateChatServerRequest(
          network_key = network_key,
          chat_room = chat_room,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: CreateChatServerRequest): CreateChatServerRequest = value.copy(
        chat_room = value.chat_room?.let(ChatRoom.ADAPTER::redact),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}
