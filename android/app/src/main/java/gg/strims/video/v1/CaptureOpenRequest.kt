// Code generated by Wire protocol buffer compiler, do not edit.
// Source: strims.video.v1.CaptureOpenRequest in video/v1/capture.proto
package gg.strims.video.v1

import com.squareup.wire.FieldEncoding
import com.squareup.wire.Message
import com.squareup.wire.ProtoAdapter
import com.squareup.wire.ProtoReader
import com.squareup.wire.ProtoWriter
import com.squareup.wire.Syntax.PROTO_3
import com.squareup.wire.WireField
import com.squareup.wire.internal.immutableCopyOf
import com.squareup.wire.internal.sanitize
import gg.strims.network.v1.DirectoryListingSnippet
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
import kotlin.hashCode
import kotlin.jvm.JvmField
import okio.ByteString

class CaptureOpenRequest(
  @field:WireField(
    tag = 1,
    adapter = "gg.strims.network.v1.DirectoryListingSnippet#ADAPTER",
    label = WireField.Label.OMIT_IDENTITY,
    jsonName = "directorySnippet"
  )
  val directory_snippet: DirectoryListingSnippet? = null,
  @field:WireField(
    tag = 2,
    adapter = "com.squareup.wire.ProtoAdapter#STRING",
    label = WireField.Label.OMIT_IDENTITY,
    jsonName = "mimeType"
  )
  val mime_type: String = "",
  network_keys: List<ByteString> = emptyList(),
  unknownFields: ByteString = ByteString.EMPTY
) : Message<CaptureOpenRequest, Nothing>(ADAPTER, unknownFields) {
  @field:WireField(
    tag = 3,
    adapter = "com.squareup.wire.ProtoAdapter#BYTES",
    label = WireField.Label.REPEATED,
    jsonName = "networkKeys"
  )
  val network_keys: List<ByteString> = immutableCopyOf("network_keys", network_keys)

  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is CaptureOpenRequest) return false
    if (unknownFields != other.unknownFields) return false
    if (directory_snippet != other.directory_snippet) return false
    if (mime_type != other.mime_type) return false
    if (network_keys != other.network_keys) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + directory_snippet.hashCode()
      result = result * 37 + mime_type.hashCode()
      result = result * 37 + network_keys.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    if (directory_snippet != null) result += """directory_snippet=$directory_snippet"""
    result += """mime_type=${sanitize(mime_type)}"""
    if (network_keys.isNotEmpty()) result += """network_keys=$network_keys"""
    return result.joinToString(prefix = "CaptureOpenRequest{", separator = ", ", postfix = "}")
  }

  fun copy(
    directory_snippet: DirectoryListingSnippet? = this.directory_snippet,
    mime_type: String = this.mime_type,
    network_keys: List<ByteString> = this.network_keys,
    unknownFields: ByteString = this.unknownFields
  ): CaptureOpenRequest = CaptureOpenRequest(directory_snippet, mime_type, network_keys,
      unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<CaptureOpenRequest> = object : ProtoAdapter<CaptureOpenRequest>(
      FieldEncoding.LENGTH_DELIMITED, 
      CaptureOpenRequest::class, 
      "type.googleapis.com/strims.video.v1.CaptureOpenRequest", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: CaptureOpenRequest): Int {
        var size = value.unknownFields.size
        if (value.directory_snippet != null) size +=
            DirectoryListingSnippet.ADAPTER.encodedSizeWithTag(1, value.directory_snippet)
        if (value.mime_type != "") size += ProtoAdapter.STRING.encodedSizeWithTag(2,
            value.mime_type)
        size += ProtoAdapter.BYTES.asRepeated().encodedSizeWithTag(3, value.network_keys)
        return size
      }

      override fun encode(writer: ProtoWriter, value: CaptureOpenRequest) {
        if (value.directory_snippet != null) DirectoryListingSnippet.ADAPTER.encodeWithTag(writer,
            1, value.directory_snippet)
        if (value.mime_type != "") ProtoAdapter.STRING.encodeWithTag(writer, 2, value.mime_type)
        ProtoAdapter.BYTES.asRepeated().encodeWithTag(writer, 3, value.network_keys)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): CaptureOpenRequest {
        var directory_snippet: DirectoryListingSnippet? = null
        var mime_type: String = ""
        val network_keys = mutableListOf<ByteString>()
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> directory_snippet = DirectoryListingSnippet.ADAPTER.decode(reader)
            2 -> mime_type = ProtoAdapter.STRING.decode(reader)
            3 -> network_keys.add(ProtoAdapter.BYTES.decode(reader))
            else -> reader.readUnknownField(tag)
          }
        }
        return CaptureOpenRequest(
          directory_snippet = directory_snippet,
          mime_type = mime_type,
          network_keys = network_keys,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: CaptureOpenRequest): CaptureOpenRequest = value.copy(
        directory_snippet = value.directory_snippet?.let(DirectoryListingSnippet.ADAPTER::redact),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}
