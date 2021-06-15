// Code generated by Wire protocol buffer compiler, do not edit.
// Source: strims.network.v1.NetworkPeerOpenRequest in network/v1/peer.proto
package gg.strims.network.v1

import com.squareup.wire.FieldEncoding
import com.squareup.wire.Message
import com.squareup.wire.ProtoAdapter
import com.squareup.wire.ProtoReader
import com.squareup.wire.ProtoWriter
import com.squareup.wire.Syntax.PROTO_3
import com.squareup.wire.WireField
import com.squareup.wire.internal.immutableCopyOf
import com.squareup.wire.internal.redactElements
import kotlin.Any
import kotlin.AssertionError
import kotlin.Boolean
import kotlin.Deprecated
import kotlin.DeprecationLevel
import kotlin.Int
import kotlin.Long
import kotlin.Nothing
import kotlin.String
import kotlin.collections.List
import kotlin.jvm.JvmField
import okio.ByteString

class NetworkPeerOpenRequest(
  bindings: List<NetworkPeerBinding> = emptyList(),
  unknownFields: ByteString = ByteString.EMPTY
) : Message<NetworkPeerOpenRequest, Nothing>(ADAPTER, unknownFields) {
  @field:WireField(
    tag = 1,
    adapter = "gg.strims.network.v1.NetworkPeerBinding#ADAPTER",
    label = WireField.Label.REPEATED
  )
  val bindings: List<NetworkPeerBinding> = immutableCopyOf("bindings", bindings)

  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is NetworkPeerOpenRequest) return false
    if (unknownFields != other.unknownFields) return false
    if (bindings != other.bindings) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + bindings.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    if (bindings.isNotEmpty()) result += """bindings=$bindings"""
    return result.joinToString(prefix = "NetworkPeerOpenRequest{", separator = ", ", postfix = "}")
  }

  fun copy(bindings: List<NetworkPeerBinding> = this.bindings, unknownFields: ByteString =
      this.unknownFields): NetworkPeerOpenRequest = NetworkPeerOpenRequest(bindings, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<NetworkPeerOpenRequest> = object :
        ProtoAdapter<NetworkPeerOpenRequest>(
      FieldEncoding.LENGTH_DELIMITED, 
      NetworkPeerOpenRequest::class, 
      "type.googleapis.com/strims.network.v1.NetworkPeerOpenRequest", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: NetworkPeerOpenRequest): Int {
        var size = value.unknownFields.size
        size += NetworkPeerBinding.ADAPTER.asRepeated().encodedSizeWithTag(1, value.bindings)
        return size
      }

      override fun encode(writer: ProtoWriter, value: NetworkPeerOpenRequest) {
        NetworkPeerBinding.ADAPTER.asRepeated().encodeWithTag(writer, 1, value.bindings)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): NetworkPeerOpenRequest {
        val bindings = mutableListOf<NetworkPeerBinding>()
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> bindings.add(NetworkPeerBinding.ADAPTER.decode(reader))
            else -> reader.readUnknownField(tag)
          }
        }
        return NetworkPeerOpenRequest(
          bindings = bindings,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: NetworkPeerOpenRequest): NetworkPeerOpenRequest = value.copy(
        bindings = value.bindings.redactElements(NetworkPeerBinding.ADAPTER),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}