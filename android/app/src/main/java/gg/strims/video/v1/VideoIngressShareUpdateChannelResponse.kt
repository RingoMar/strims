// Code generated by Wire protocol buffer compiler, do not edit.
// Source: strims.video.v1.VideoIngressShareUpdateChannelResponse in video/v1/ingress.proto
package gg.strims.video.v1

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

class VideoIngressShareUpdateChannelResponse(
  @field:WireField(
    tag = 1,
    adapter = "gg.strims.video.v1.VideoChannel#ADAPTER",
    label = WireField.Label.OMIT_IDENTITY
  )
  val channel: VideoChannel? = null,
  unknownFields: ByteString = ByteString.EMPTY
) : Message<VideoIngressShareUpdateChannelResponse, Nothing>(ADAPTER, unknownFields) {
  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is VideoIngressShareUpdateChannelResponse) return false
    if (unknownFields != other.unknownFields) return false
    if (channel != other.channel) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + channel.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    if (channel != null) result += """channel=$channel"""
    return result.joinToString(prefix = "VideoIngressShareUpdateChannelResponse{", separator = ", ",
        postfix = "}")
  }

  fun copy(channel: VideoChannel? = this.channel, unknownFields: ByteString = this.unknownFields):
      VideoIngressShareUpdateChannelResponse = VideoIngressShareUpdateChannelResponse(channel,
      unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<VideoIngressShareUpdateChannelResponse> = object :
        ProtoAdapter<VideoIngressShareUpdateChannelResponse>(
      FieldEncoding.LENGTH_DELIMITED, 
      VideoIngressShareUpdateChannelResponse::class, 
      "type.googleapis.com/strims.video.v1.VideoIngressShareUpdateChannelResponse", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: VideoIngressShareUpdateChannelResponse): Int {
        var size = value.unknownFields.size
        if (value.channel != null) size += VideoChannel.ADAPTER.encodedSizeWithTag(1, value.channel)
        return size
      }

      override fun encode(writer: ProtoWriter, value: VideoIngressShareUpdateChannelResponse) {
        if (value.channel != null) VideoChannel.ADAPTER.encodeWithTag(writer, 1, value.channel)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): VideoIngressShareUpdateChannelResponse {
        var channel: VideoChannel? = null
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> channel = VideoChannel.ADAPTER.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return VideoIngressShareUpdateChannelResponse(
          channel = channel,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: VideoIngressShareUpdateChannelResponse):
          VideoIngressShareUpdateChannelResponse = value.copy(
        channel = value.channel?.let(VideoChannel.ADAPTER::redact),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}