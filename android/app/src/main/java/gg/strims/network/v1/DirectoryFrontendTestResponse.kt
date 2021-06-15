// Code generated by Wire protocol buffer compiler, do not edit.
// Source: strims.network.v1.DirectoryFrontendTestResponse in network/v1/directory.proto
package gg.strims.network.v1

import com.squareup.wire.FieldEncoding
import com.squareup.wire.Message
import com.squareup.wire.ProtoAdapter
import com.squareup.wire.ProtoReader
import com.squareup.wire.ProtoWriter
import com.squareup.wire.Syntax.PROTO_3
import kotlin.Any
import kotlin.AssertionError
import kotlin.Boolean
import kotlin.Deprecated
import kotlin.DeprecationLevel
import kotlin.Int
import kotlin.Long
import kotlin.Nothing
import kotlin.String
import kotlin.jvm.JvmField
import okio.ByteString

class DirectoryFrontendTestResponse(
  unknownFields: ByteString = ByteString.EMPTY
) : Message<DirectoryFrontendTestResponse, Nothing>(ADAPTER, unknownFields) {
  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is DirectoryFrontendTestResponse) return false
    if (unknownFields != other.unknownFields) return false
    return true
  }

  override fun hashCode(): Int = unknownFields.hashCode()

  override fun toString(): String = "DirectoryFrontendTestResponse{}"

  fun copy(unknownFields: ByteString = this.unknownFields): DirectoryFrontendTestResponse =
      DirectoryFrontendTestResponse(unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<DirectoryFrontendTestResponse> = object :
        ProtoAdapter<DirectoryFrontendTestResponse>(
      FieldEncoding.LENGTH_DELIMITED, 
      DirectoryFrontendTestResponse::class, 
      "type.googleapis.com/strims.network.v1.DirectoryFrontendTestResponse", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: DirectoryFrontendTestResponse): Int {
        var size = value.unknownFields.size
        return size
      }

      override fun encode(writer: ProtoWriter, value: DirectoryFrontendTestResponse) {
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): DirectoryFrontendTestResponse {
        val unknownFields = reader.forEachTag(reader::readUnknownField)
        return DirectoryFrontendTestResponse(
          unknownFields = unknownFields
        )
      }

      override fun redact(value: DirectoryFrontendTestResponse): DirectoryFrontendTestResponse =
          value.copy(
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}