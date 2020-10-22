// Code generated by Wire protocol buffer compiler, do not edit.
// Source: OpenDirectoryClientRequest in directory.proto
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

class OpenDirectoryClientRequest(
  @field:WireField(
    tag = 1,
    adapter = "com.squareup.wire.ProtoAdapter#BYTES",
    label = WireField.Label.OMIT_IDENTITY,
    jsonName = "networkKey"
  )
  val network_key: ByteString = ByteString.EMPTY,
  @field:WireField(
    tag = 2,
    adapter = "com.squareup.wire.ProtoAdapter#BYTES",
    label = WireField.Label.OMIT_IDENTITY,
    jsonName = "serverKey"
  )
  val server_key: ByteString = ByteString.EMPTY,
  unknownFields: ByteString = ByteString.EMPTY
) : Message<OpenDirectoryClientRequest, Nothing>(ADAPTER, unknownFields) {
  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is OpenDirectoryClientRequest) return false
    if (unknownFields != other.unknownFields) return false
    if (network_key != other.network_key) return false
    if (server_key != other.server_key) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + network_key.hashCode()
      result = result * 37 + server_key.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    result += """network_key=$network_key"""
    result += """server_key=$server_key"""
    return result.joinToString(prefix = "OpenDirectoryClientRequest{", separator = ", ", postfix =
        "}")
  }

  fun copy(
    network_key: ByteString = this.network_key,
    server_key: ByteString = this.server_key,
    unknownFields: ByteString = this.unknownFields
  ): OpenDirectoryClientRequest = OpenDirectoryClientRequest(network_key, server_key, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<OpenDirectoryClientRequest> = object :
        ProtoAdapter<OpenDirectoryClientRequest>(
      FieldEncoding.LENGTH_DELIMITED, 
      OpenDirectoryClientRequest::class, 
      "type.googleapis.com/OpenDirectoryClientRequest", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: OpenDirectoryClientRequest): Int {
        var size = value.unknownFields.size
        if (value.network_key != ByteString.EMPTY) size += ProtoAdapter.BYTES.encodedSizeWithTag(1,
            value.network_key)
        if (value.server_key != ByteString.EMPTY) size += ProtoAdapter.BYTES.encodedSizeWithTag(2,
            value.server_key)
        return size
      }

      override fun encode(writer: ProtoWriter, value: OpenDirectoryClientRequest) {
        if (value.network_key != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 1,
            value.network_key)
        if (value.server_key != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 2,
            value.server_key)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): OpenDirectoryClientRequest {
        var network_key: ByteString = ByteString.EMPTY
        var server_key: ByteString = ByteString.EMPTY
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> network_key = ProtoAdapter.BYTES.decode(reader)
            2 -> server_key = ProtoAdapter.BYTES.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return OpenDirectoryClientRequest(
          network_key = network_key,
          server_key = server_key,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: OpenDirectoryClientRequest): OpenDirectoryClientRequest =
          value.copy(
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}