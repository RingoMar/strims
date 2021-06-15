// Code generated by Wire protocol buffer compiler, do not edit.
// Source: strims.profile.v1.StorageKey in profile/v1/profile.proto
package gg.strims.profile.v1

import com.squareup.wire.FieldEncoding
import com.squareup.wire.Message
import com.squareup.wire.ProtoAdapter
import com.squareup.wire.ProtoReader
import com.squareup.wire.ProtoWriter
import com.squareup.wire.Syntax
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

class StorageKey(
  @field:WireField(
    tag = 1,
    adapter = "gg.strims.profile.v1.KDFType#ADAPTER",
    label = WireField.Label.OMIT_IDENTITY,
    jsonName = "kdfType"
  )
  val kdf_type: KDFType = KDFType.KDF_TYPE_UNDEFINED,
  @field:WireField(
    tag = 2,
    adapter = "gg.strims.profile.v1.StorageKey${'$'}PBKDF2Options#ADAPTER",
    jsonName = "pbkdf2Options"
  )
  val pbkdf2_options: PBKDF2Options? = null,
  unknownFields: ByteString = ByteString.EMPTY
) : Message<StorageKey, Nothing>(ADAPTER, unknownFields) {
  init {
  }

  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is StorageKey) return false
    if (unknownFields != other.unknownFields) return false
    if (kdf_type != other.kdf_type) return false
    if (pbkdf2_options != other.pbkdf2_options) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + kdf_type.hashCode()
      result = result * 37 + pbkdf2_options.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    result += """kdf_type=$kdf_type"""
    if (pbkdf2_options != null) result += """pbkdf2_options=$pbkdf2_options"""
    return result.joinToString(prefix = "StorageKey{", separator = ", ", postfix = "}")
  }

  fun copy(
    kdf_type: KDFType = this.kdf_type,
    pbkdf2_options: PBKDF2Options? = this.pbkdf2_options,
    unknownFields: ByteString = this.unknownFields
  ): StorageKey = StorageKey(kdf_type, pbkdf2_options, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<StorageKey> = object : ProtoAdapter<StorageKey>(
      FieldEncoding.LENGTH_DELIMITED, 
      StorageKey::class, 
      "type.googleapis.com/strims.profile.v1.StorageKey", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: StorageKey): Int {
        var size = value.unknownFields.size
        if (value.kdf_type != KDFType.KDF_TYPE_UNDEFINED) size +=
            KDFType.ADAPTER.encodedSizeWithTag(1, value.kdf_type)
        size += PBKDF2Options.ADAPTER.encodedSizeWithTag(2, value.pbkdf2_options)
        return size
      }

      override fun encode(writer: ProtoWriter, value: StorageKey) {
        if (value.kdf_type != KDFType.KDF_TYPE_UNDEFINED) KDFType.ADAPTER.encodeWithTag(writer, 1,
            value.kdf_type)
        PBKDF2Options.ADAPTER.encodeWithTag(writer, 2, value.pbkdf2_options)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): StorageKey {
        var kdf_type: KDFType = KDFType.KDF_TYPE_UNDEFINED
        var pbkdf2_options: PBKDF2Options? = null
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> try {
              kdf_type = KDFType.ADAPTER.decode(reader)
            } catch (e: ProtoAdapter.EnumConstantNotFoundException) {
              reader.addUnknownField(tag, FieldEncoding.VARINT, e.value.toLong())
            }
            2 -> pbkdf2_options = PBKDF2Options.ADAPTER.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return StorageKey(
          kdf_type = kdf_type,
          pbkdf2_options = pbkdf2_options,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: StorageKey): StorageKey = value.copy(
        pbkdf2_options = value.pbkdf2_options?.let(PBKDF2Options.ADAPTER::redact),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }

  class PBKDF2Options(
    @field:WireField(
      tag = 1,
      adapter = "com.squareup.wire.ProtoAdapter#UINT32",
      label = WireField.Label.OMIT_IDENTITY
    )
    val iterations: Int = 0,
    @field:WireField(
      tag = 2,
      adapter = "com.squareup.wire.ProtoAdapter#UINT32",
      label = WireField.Label.OMIT_IDENTITY,
      jsonName = "keySize"
    )
    val key_size: Int = 0,
    @field:WireField(
      tag = 3,
      adapter = "com.squareup.wire.ProtoAdapter#BYTES",
      label = WireField.Label.OMIT_IDENTITY
    )
    val salt: ByteString = ByteString.EMPTY,
    unknownFields: ByteString = ByteString.EMPTY
  ) : Message<PBKDF2Options, Nothing>(ADAPTER, unknownFields) {
    @Deprecated(
      message = "Shouldn't be used in Kotlin",
      level = DeprecationLevel.HIDDEN
    )
    override fun newBuilder(): Nothing = throw AssertionError()

    override fun equals(other: Any?): Boolean {
      if (other === this) return true
      if (other !is PBKDF2Options) return false
      if (unknownFields != other.unknownFields) return false
      if (iterations != other.iterations) return false
      if (key_size != other.key_size) return false
      if (salt != other.salt) return false
      return true
    }

    override fun hashCode(): Int {
      var result = super.hashCode
      if (result == 0) {
        result = unknownFields.hashCode()
        result = result * 37 + iterations.hashCode()
        result = result * 37 + key_size.hashCode()
        result = result * 37 + salt.hashCode()
        super.hashCode = result
      }
      return result
    }

    override fun toString(): String {
      val result = mutableListOf<String>()
      result += """iterations=$iterations"""
      result += """key_size=$key_size"""
      result += """salt=$salt"""
      return result.joinToString(prefix = "PBKDF2Options{", separator = ", ", postfix = "}")
    }

    fun copy(
      iterations: Int = this.iterations,
      key_size: Int = this.key_size,
      salt: ByteString = this.salt,
      unknownFields: ByteString = this.unknownFields
    ): PBKDF2Options = PBKDF2Options(iterations, key_size, salt, unknownFields)

    companion object {
      @JvmField
      val ADAPTER: ProtoAdapter<PBKDF2Options> = object : ProtoAdapter<PBKDF2Options>(
        FieldEncoding.LENGTH_DELIMITED, 
        PBKDF2Options::class, 
        "type.googleapis.com/strims.profile.v1.StorageKey.PBKDF2Options", 
        PROTO_3, 
        null
      ) {
        override fun encodedSize(value: PBKDF2Options): Int {
          var size = value.unknownFields.size
          if (value.iterations != 0) size += ProtoAdapter.UINT32.encodedSizeWithTag(1,
              value.iterations)
          if (value.key_size != 0) size += ProtoAdapter.UINT32.encodedSizeWithTag(2, value.key_size)
          if (value.salt != ByteString.EMPTY) size += ProtoAdapter.BYTES.encodedSizeWithTag(3,
              value.salt)
          return size
        }

        override fun encode(writer: ProtoWriter, value: PBKDF2Options) {
          if (value.iterations != 0) ProtoAdapter.UINT32.encodeWithTag(writer, 1, value.iterations)
          if (value.key_size != 0) ProtoAdapter.UINT32.encodeWithTag(writer, 2, value.key_size)
          if (value.salt != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 3,
              value.salt)
          writer.writeBytes(value.unknownFields)
        }

        override fun decode(reader: ProtoReader): PBKDF2Options {
          var iterations: Int = 0
          var key_size: Int = 0
          var salt: ByteString = ByteString.EMPTY
          val unknownFields = reader.forEachTag { tag ->
            when (tag) {
              1 -> iterations = ProtoAdapter.UINT32.decode(reader)
              2 -> key_size = ProtoAdapter.UINT32.decode(reader)
              3 -> salt = ProtoAdapter.BYTES.decode(reader)
              else -> reader.readUnknownField(tag)
            }
          }
          return PBKDF2Options(
            iterations = iterations,
            key_size = key_size,
            salt = salt,
            unknownFields = unknownFields
          )
        }

        override fun redact(value: PBKDF2Options): PBKDF2Options = value.copy(
          unknownFields = ByteString.EMPTY
        )
      }

      private const val serialVersionUID: Long = 0L
    }
  }
}