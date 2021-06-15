// Code generated by Wire protocol buffer compiler, do not edit.
// Source: strims.transfer.v1.Transfer in transfer/v1/transfer.proto
package gg.strims.transfer.v1

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

class Transfer(
  @field:WireField(
    tag = 1,
    adapter = "com.squareup.wire.ProtoAdapter#BYTES",
    label = WireField.Label.OMIT_IDENTITY
  )
  val id: ByteString = ByteString.EMPTY,
  unknownFields: ByteString = ByteString.EMPTY
) : Message<Transfer, Nothing>(ADAPTER, unknownFields) {
  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is Transfer) return false
    if (unknownFields != other.unknownFields) return false
    if (id != other.id) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + id.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    result += """id=$id"""
    return result.joinToString(prefix = "Transfer{", separator = ", ", postfix = "}")
  }

  fun copy(id: ByteString = this.id, unknownFields: ByteString = this.unknownFields): Transfer =
      Transfer(id, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<Transfer> = object : ProtoAdapter<Transfer>(
      FieldEncoding.LENGTH_DELIMITED, 
      Transfer::class, 
      "type.googleapis.com/strims.transfer.v1.Transfer", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: Transfer): Int {
        var size = value.unknownFields.size
        if (value.id != ByteString.EMPTY) size += ProtoAdapter.BYTES.encodedSizeWithTag(1, value.id)
        return size
      }

      override fun encode(writer: ProtoWriter, value: Transfer) {
        if (value.id != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 1, value.id)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): Transfer {
        var id: ByteString = ByteString.EMPTY
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> id = ProtoAdapter.BYTES.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return Transfer(
          id = id,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: Transfer): Transfer = value.copy(
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}