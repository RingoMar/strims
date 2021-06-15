// Code generated by Wire protocol buffer compiler, do not edit.
// Source: strims.network.v1.DirectoryListing in network/v1/directory.proto
package gg.strims.network.v1

import com.squareup.wire.FieldEncoding
import com.squareup.wire.Message
import com.squareup.wire.ProtoAdapter
import com.squareup.wire.ProtoReader
import com.squareup.wire.ProtoWriter
import com.squareup.wire.Syntax.PROTO_3
import com.squareup.wire.WireField
import com.squareup.wire.internal.countNonNull
import gg.strims.type.Certificate
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

class DirectoryListing(
  @field:WireField(
    tag = 1,
    adapter = "gg.strims.type.Certificate#ADAPTER",
    label = WireField.Label.OMIT_IDENTITY
  )
  val creator: Certificate? = null,
  @field:WireField(
    tag = 2,
    adapter = "com.squareup.wire.ProtoAdapter#INT64",
    label = WireField.Label.OMIT_IDENTITY
  )
  val timestamp: Long = 0L,
  @field:WireField(
    tag = 3,
    adapter = "gg.strims.network.v1.DirectoryListingSnippet#ADAPTER",
    label = WireField.Label.OMIT_IDENTITY
  )
  val snippet: DirectoryListingSnippet? = null,
  @field:WireField(
    tag = 10001,
    adapter = "com.squareup.wire.ProtoAdapter#BYTES",
    label = WireField.Label.OMIT_IDENTITY
  )
  val key: ByteString = ByteString.EMPTY,
  @field:WireField(
    tag = 10002,
    adapter = "com.squareup.wire.ProtoAdapter#BYTES",
    label = WireField.Label.OMIT_IDENTITY
  )
  val signature: ByteString = ByteString.EMPTY,
  @field:WireField(
    tag = 1001,
    adapter = "gg.strims.network.v1.DirectoryListingMedia#ADAPTER"
  )
  val media: DirectoryListingMedia? = null,
  @field:WireField(
    tag = 1002,
    adapter = "gg.strims.network.v1.DirectoryListingService#ADAPTER"
  )
  val service: DirectoryListingService? = null,
  unknownFields: ByteString = ByteString.EMPTY
) : Message<DirectoryListing, Nothing>(ADAPTER, unknownFields) {
  init {
    require(countNonNull(media, service) <= 1) {
      "At most one of media, service may be non-null"
    }
  }

  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is DirectoryListing) return false
    if (unknownFields != other.unknownFields) return false
    if (creator != other.creator) return false
    if (timestamp != other.timestamp) return false
    if (snippet != other.snippet) return false
    if (key != other.key) return false
    if (signature != other.signature) return false
    if (media != other.media) return false
    if (service != other.service) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + creator.hashCode()
      result = result * 37 + timestamp.hashCode()
      result = result * 37 + snippet.hashCode()
      result = result * 37 + key.hashCode()
      result = result * 37 + signature.hashCode()
      result = result * 37 + media.hashCode()
      result = result * 37 + service.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    if (creator != null) result += """creator=$creator"""
    result += """timestamp=$timestamp"""
    if (snippet != null) result += """snippet=$snippet"""
    result += """key=$key"""
    result += """signature=$signature"""
    if (media != null) result += """media=$media"""
    if (service != null) result += """service=$service"""
    return result.joinToString(prefix = "DirectoryListing{", separator = ", ", postfix = "}")
  }

  fun copy(
    creator: Certificate? = this.creator,
    timestamp: Long = this.timestamp,
    snippet: DirectoryListingSnippet? = this.snippet,
    key: ByteString = this.key,
    signature: ByteString = this.signature,
    media: DirectoryListingMedia? = this.media,
    service: DirectoryListingService? = this.service,
    unknownFields: ByteString = this.unknownFields
  ): DirectoryListing = DirectoryListing(creator, timestamp, snippet, key, signature, media,
      service, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<DirectoryListing> = object : ProtoAdapter<DirectoryListing>(
      FieldEncoding.LENGTH_DELIMITED, 
      DirectoryListing::class, 
      "type.googleapis.com/strims.network.v1.DirectoryListing", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: DirectoryListing): Int {
        var size = value.unknownFields.size
        if (value.creator != null) size += Certificate.ADAPTER.encodedSizeWithTag(1, value.creator)
        if (value.timestamp != 0L) size += ProtoAdapter.INT64.encodedSizeWithTag(2, value.timestamp)
        if (value.snippet != null) size += DirectoryListingSnippet.ADAPTER.encodedSizeWithTag(3,
            value.snippet)
        if (value.key != ByteString.EMPTY) size += ProtoAdapter.BYTES.encodedSizeWithTag(10001,
            value.key)
        if (value.signature != ByteString.EMPTY) size +=
            ProtoAdapter.BYTES.encodedSizeWithTag(10002, value.signature)
        size += DirectoryListingMedia.ADAPTER.encodedSizeWithTag(1001, value.media)
        size += DirectoryListingService.ADAPTER.encodedSizeWithTag(1002, value.service)
        return size
      }

      override fun encode(writer: ProtoWriter, value: DirectoryListing) {
        if (value.creator != null) Certificate.ADAPTER.encodeWithTag(writer, 1, value.creator)
        if (value.timestamp != 0L) ProtoAdapter.INT64.encodeWithTag(writer, 2, value.timestamp)
        if (value.snippet != null) DirectoryListingSnippet.ADAPTER.encodeWithTag(writer, 3,
            value.snippet)
        if (value.key != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 10001,
            value.key)
        if (value.signature != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 10002,
            value.signature)
        DirectoryListingMedia.ADAPTER.encodeWithTag(writer, 1001, value.media)
        DirectoryListingService.ADAPTER.encodeWithTag(writer, 1002, value.service)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): DirectoryListing {
        var creator: Certificate? = null
        var timestamp: Long = 0L
        var snippet: DirectoryListingSnippet? = null
        var key: ByteString = ByteString.EMPTY
        var signature: ByteString = ByteString.EMPTY
        var media: DirectoryListingMedia? = null
        var service: DirectoryListingService? = null
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> creator = Certificate.ADAPTER.decode(reader)
            2 -> timestamp = ProtoAdapter.INT64.decode(reader)
            3 -> snippet = DirectoryListingSnippet.ADAPTER.decode(reader)
            10001 -> key = ProtoAdapter.BYTES.decode(reader)
            10002 -> signature = ProtoAdapter.BYTES.decode(reader)
            1001 -> media = DirectoryListingMedia.ADAPTER.decode(reader)
            1002 -> service = DirectoryListingService.ADAPTER.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return DirectoryListing(
          creator = creator,
          timestamp = timestamp,
          snippet = snippet,
          key = key,
          signature = signature,
          media = media,
          service = service,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: DirectoryListing): DirectoryListing = value.copy(
        creator = value.creator?.let(Certificate.ADAPTER::redact),
        snippet = value.snippet?.let(DirectoryListingSnippet.ADAPTER::redact),
        media = value.media?.let(DirectoryListingMedia.ADAPTER::redact),
        service = value.service?.let(DirectoryListingService.ADAPTER::redact),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}