// Code generated by Wire protocol buffer compiler, do not edit.
// Source: strims.video.v1.VideoIngressShareUpdateChannelRequest in video/v1/ingress.proto
package gg.strims.video.v1

import com.squareup.wire.FieldEncoding
import com.squareup.wire.Message
import com.squareup.wire.ProtoAdapter
import com.squareup.wire.ProtoReader
import com.squareup.wire.ProtoWriter
import com.squareup.wire.Syntax.PROTO_3
import com.squareup.wire.WireField
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
import kotlin.hashCode
import kotlin.jvm.JvmField
import okio.ByteString

class VideoIngressShareUpdateChannelRequest(
  @field:WireField(
    tag = 1,
    adapter = "gg.strims.network.v1.DirectoryListingSnippet#ADAPTER",
    label = WireField.Label.OMIT_IDENTITY,
    jsonName = "directoryListingSnippet"
  )
  val directory_listing_snippet: DirectoryListingSnippet? = null,
  unknownFields: ByteString = ByteString.EMPTY
) : Message<VideoIngressShareUpdateChannelRequest, Nothing>(ADAPTER, unknownFields) {
  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is VideoIngressShareUpdateChannelRequest) return false
    if (unknownFields != other.unknownFields) return false
    if (directory_listing_snippet != other.directory_listing_snippet) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + directory_listing_snippet.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    if (directory_listing_snippet != null) result +=
        """directory_listing_snippet=$directory_listing_snippet"""
    return result.joinToString(prefix = "VideoIngressShareUpdateChannelRequest{", separator = ", ",
        postfix = "}")
  }

  fun copy(directory_listing_snippet: DirectoryListingSnippet? = this.directory_listing_snippet,
      unknownFields: ByteString = this.unknownFields): VideoIngressShareUpdateChannelRequest =
      VideoIngressShareUpdateChannelRequest(directory_listing_snippet, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<VideoIngressShareUpdateChannelRequest> = object :
        ProtoAdapter<VideoIngressShareUpdateChannelRequest>(
      FieldEncoding.LENGTH_DELIMITED, 
      VideoIngressShareUpdateChannelRequest::class, 
      "type.googleapis.com/strims.video.v1.VideoIngressShareUpdateChannelRequest", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: VideoIngressShareUpdateChannelRequest): Int {
        var size = value.unknownFields.size
        if (value.directory_listing_snippet != null) size +=
            DirectoryListingSnippet.ADAPTER.encodedSizeWithTag(1, value.directory_listing_snippet)
        return size
      }

      override fun encode(writer: ProtoWriter, value: VideoIngressShareUpdateChannelRequest) {
        if (value.directory_listing_snippet != null)
            DirectoryListingSnippet.ADAPTER.encodeWithTag(writer, 1,
            value.directory_listing_snippet)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): VideoIngressShareUpdateChannelRequest {
        var directory_listing_snippet: DirectoryListingSnippet? = null
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> directory_listing_snippet = DirectoryListingSnippet.ADAPTER.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return VideoIngressShareUpdateChannelRequest(
          directory_listing_snippet = directory_listing_snippet,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: VideoIngressShareUpdateChannelRequest):
          VideoIngressShareUpdateChannelRequest = value.copy(
        directory_listing_snippet =
            value.directory_listing_snippet?.let(DirectoryListingSnippet.ADAPTER::redact),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}