// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: protocol/burst.proto

package burst.protocol;

/**
 * Protobuf type {@code protocol.BurstMessage}
 */
public final class BurstMessage extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:protocol.BurstMessage)
    BurstMessageOrBuilder {
private static final long serialVersionUID = 0L;
  // Use BurstMessage.newBuilder() to construct.
  private BurstMessage(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private BurstMessage() {
    type_ = 0;
    data_ = com.google.protobuf.ByteString.EMPTY;
    serverPort_ = emptyIntList();
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new BurstMessage();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private BurstMessage(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 8: {
            int rawValue = input.readEnum();

            type_ = rawValue;
            break;
          }
          case 18: {
            if (!((mutable_bitField0_ & 0x00000001) != 0)) {
              header_ = com.google.protobuf.MapField.newMapField(
                  HeaderDefaultEntryHolder.defaultEntry);
              mutable_bitField0_ |= 0x00000001;
            }
            com.google.protobuf.MapEntry<java.lang.Integer, com.google.protobuf.Any>
            header__ = input.readMessage(
                HeaderDefaultEntryHolder.defaultEntry.getParserForType(), extensionRegistry);
            header_.getMutableMap().put(
                header__.getKey(), header__.getValue());
            break;
          }
          case 26: {

            data_ = input.readBytes();
            break;
          }
          case 32: {
            if (!((mutable_bitField0_ & 0x00000002) != 0)) {
              serverPort_ = newIntList();
              mutable_bitField0_ |= 0x00000002;
            }
            serverPort_.addInt(input.readInt32());
            break;
          }
          case 34: {
            int length = input.readRawVarint32();
            int limit = input.pushLimit(length);
            if (!((mutable_bitField0_ & 0x00000002) != 0) && input.getBytesUntilLimit() > 0) {
              serverPort_ = newIntList();
              mutable_bitField0_ |= 0x00000002;
            }
            while (input.getBytesUntilLimit() > 0) {
              serverPort_.addInt(input.readInt32());
            }
            input.popLimit(limit);
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      if (((mutable_bitField0_ & 0x00000002) != 0)) {
        serverPort_.makeImmutable(); // C
      }
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return burst.protocol.Burst.internal_static_protocol_BurstMessage_descriptor;
  }

  @SuppressWarnings({"rawtypes"})
  @java.lang.Override
  protected com.google.protobuf.MapField internalGetMapField(
      int number) {
    switch (number) {
      case 2:
        return internalGetHeader();
      default:
        throw new RuntimeException(
            "Invalid map field number: " + number);
    }
  }
  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return burst.protocol.Burst.internal_static_protocol_BurstMessage_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            burst.protocol.BurstMessage.class, burst.protocol.BurstMessage.Builder.class);
  }

  public static final int TYPE_FIELD_NUMBER = 1;
  private int type_;
  /**
   * <code>.protocol.BurstType type = 1;</code>
   * @return The enum numeric value on the wire for type.
   */
  @java.lang.Override public int getTypeValue() {
    return type_;
  }
  /**
   * <code>.protocol.BurstType type = 1;</code>
   * @return The type.
   */
  @java.lang.Override public burst.protocol.BurstType getType() {
    @SuppressWarnings("deprecation")
    burst.protocol.BurstType result = burst.protocol.BurstType.valueOf(type_);
    return result == null ? burst.protocol.BurstType.UNRECOGNIZED : result;
  }

  public static final int HEADER_FIELD_NUMBER = 2;
  private static final class HeaderDefaultEntryHolder {
    static final com.google.protobuf.MapEntry<
        java.lang.Integer, com.google.protobuf.Any> defaultEntry =
            com.google.protobuf.MapEntry
            .<java.lang.Integer, com.google.protobuf.Any>newDefaultInstance(
                burst.protocol.Burst.internal_static_protocol_BurstMessage_HeaderEntry_descriptor, 
                com.google.protobuf.WireFormat.FieldType.INT32,
                0,
                com.google.protobuf.WireFormat.FieldType.MESSAGE,
                com.google.protobuf.Any.getDefaultInstance());
  }
  private com.google.protobuf.MapField<
      java.lang.Integer, com.google.protobuf.Any> header_;
  private com.google.protobuf.MapField<java.lang.Integer, com.google.protobuf.Any>
  internalGetHeader() {
    if (header_ == null) {
      return com.google.protobuf.MapField.emptyMapField(
          HeaderDefaultEntryHolder.defaultEntry);
    }
    return header_;
  }

  public int getHeaderCount() {
    return internalGetHeader().getMap().size();
  }
  /**
   * <code>map&lt;int32, .google.protobuf.Any&gt; header = 2;</code>
   */

  @java.lang.Override
  public boolean containsHeader(
      int key) {
    
    return internalGetHeader().getMap().containsKey(key);
  }
  /**
   * Use {@link #getHeaderMap()} instead.
   */
  @java.lang.Override
  @java.lang.Deprecated
  public java.util.Map<java.lang.Integer, com.google.protobuf.Any> getHeader() {
    return getHeaderMap();
  }
  /**
   * <code>map&lt;int32, .google.protobuf.Any&gt; header = 2;</code>
   */
  @java.lang.Override

  public java.util.Map<java.lang.Integer, com.google.protobuf.Any> getHeaderMap() {
    return internalGetHeader().getMap();
  }
  /**
   * <code>map&lt;int32, .google.protobuf.Any&gt; header = 2;</code>
   */
  @java.lang.Override

  public com.google.protobuf.Any getHeaderOrDefault(
      int key,
      com.google.protobuf.Any defaultValue) {
    
    java.util.Map<java.lang.Integer, com.google.protobuf.Any> map =
        internalGetHeader().getMap();
    return map.containsKey(key) ? map.get(key) : defaultValue;
  }
  /**
   * <code>map&lt;int32, .google.protobuf.Any&gt; header = 2;</code>
   */
  @java.lang.Override

  public com.google.protobuf.Any getHeaderOrThrow(
      int key) {
    
    java.util.Map<java.lang.Integer, com.google.protobuf.Any> map =
        internalGetHeader().getMap();
    if (!map.containsKey(key)) {
      throw new java.lang.IllegalArgumentException();
    }
    return map.get(key);
  }

  public static final int DATA_FIELD_NUMBER = 3;
  private com.google.protobuf.ByteString data_;
  /**
   * <code>bytes data = 3;</code>
   * @return The data.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString getData() {
    return data_;
  }

  public static final int SERVERPORT_FIELD_NUMBER = 4;
  private com.google.protobuf.Internal.IntList serverPort_;
  /**
   * <pre>
   * 当type = REMOVE_PROXY_INFO 时有效
   * </pre>
   *
   * <code>repeated int32 serverPort = 4;</code>
   * @return A list containing the serverPort.
   */
  @java.lang.Override
  public java.util.List<java.lang.Integer>
      getServerPortList() {
    return serverPort_;
  }
  /**
   * <pre>
   * 当type = REMOVE_PROXY_INFO 时有效
   * </pre>
   *
   * <code>repeated int32 serverPort = 4;</code>
   * @return The count of serverPort.
   */
  public int getServerPortCount() {
    return serverPort_.size();
  }
  /**
   * <pre>
   * 当type = REMOVE_PROXY_INFO 时有效
   * </pre>
   *
   * <code>repeated int32 serverPort = 4;</code>
   * @param index The index of the element to return.
   * @return The serverPort at the given index.
   */
  public int getServerPort(int index) {
    return serverPort_.getInt(index);
  }
  private int serverPortMemoizedSerializedSize = -1;

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    getSerializedSize();
    if (type_ != burst.protocol.BurstType.ADD_PROXY_INFO.getNumber()) {
      output.writeEnum(1, type_);
    }
    com.google.protobuf.GeneratedMessageV3
      .serializeIntegerMapTo(
        output,
        internalGetHeader(),
        HeaderDefaultEntryHolder.defaultEntry,
        2);
    if (!data_.isEmpty()) {
      output.writeBytes(3, data_);
    }
    if (getServerPortList().size() > 0) {
      output.writeUInt32NoTag(34);
      output.writeUInt32NoTag(serverPortMemoizedSerializedSize);
    }
    for (int i = 0; i < serverPort_.size(); i++) {
      output.writeInt32NoTag(serverPort_.getInt(i));
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (type_ != burst.protocol.BurstType.ADD_PROXY_INFO.getNumber()) {
      size += com.google.protobuf.CodedOutputStream
        .computeEnumSize(1, type_);
    }
    for (java.util.Map.Entry<java.lang.Integer, com.google.protobuf.Any> entry
         : internalGetHeader().getMap().entrySet()) {
      com.google.protobuf.MapEntry<java.lang.Integer, com.google.protobuf.Any>
      header__ = HeaderDefaultEntryHolder.defaultEntry.newBuilderForType()
          .setKey(entry.getKey())
          .setValue(entry.getValue())
          .build();
      size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(2, header__);
    }
    if (!data_.isEmpty()) {
      size += com.google.protobuf.CodedOutputStream
        .computeBytesSize(3, data_);
    }
    {
      int dataSize = 0;
      for (int i = 0; i < serverPort_.size(); i++) {
        dataSize += com.google.protobuf.CodedOutputStream
          .computeInt32SizeNoTag(serverPort_.getInt(i));
      }
      size += dataSize;
      if (!getServerPortList().isEmpty()) {
        size += 1;
        size += com.google.protobuf.CodedOutputStream
            .computeInt32SizeNoTag(dataSize);
      }
      serverPortMemoizedSerializedSize = dataSize;
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof burst.protocol.BurstMessage)) {
      return super.equals(obj);
    }
    burst.protocol.BurstMessage other = (burst.protocol.BurstMessage) obj;

    if (type_ != other.type_) return false;
    if (!internalGetHeader().equals(
        other.internalGetHeader())) return false;
    if (!getData()
        .equals(other.getData())) return false;
    if (!getServerPortList()
        .equals(other.getServerPortList())) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + TYPE_FIELD_NUMBER;
    hash = (53 * hash) + type_;
    if (!internalGetHeader().getMap().isEmpty()) {
      hash = (37 * hash) + HEADER_FIELD_NUMBER;
      hash = (53 * hash) + internalGetHeader().hashCode();
    }
    hash = (37 * hash) + DATA_FIELD_NUMBER;
    hash = (53 * hash) + getData().hashCode();
    if (getServerPortCount() > 0) {
      hash = (37 * hash) + SERVERPORT_FIELD_NUMBER;
      hash = (53 * hash) + getServerPortList().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static burst.protocol.BurstMessage parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static burst.protocol.BurstMessage parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static burst.protocol.BurstMessage parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static burst.protocol.BurstMessage parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static burst.protocol.BurstMessage parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static burst.protocol.BurstMessage parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static burst.protocol.BurstMessage parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static burst.protocol.BurstMessage parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static burst.protocol.BurstMessage parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static burst.protocol.BurstMessage parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static burst.protocol.BurstMessage parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static burst.protocol.BurstMessage parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(burst.protocol.BurstMessage prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code protocol.BurstMessage}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:protocol.BurstMessage)
      burst.protocol.BurstMessageOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return burst.protocol.Burst.internal_static_protocol_BurstMessage_descriptor;
    }

    @SuppressWarnings({"rawtypes"})
    protected com.google.protobuf.MapField internalGetMapField(
        int number) {
      switch (number) {
        case 2:
          return internalGetHeader();
        default:
          throw new RuntimeException(
              "Invalid map field number: " + number);
      }
    }
    @SuppressWarnings({"rawtypes"})
    protected com.google.protobuf.MapField internalGetMutableMapField(
        int number) {
      switch (number) {
        case 2:
          return internalGetMutableHeader();
        default:
          throw new RuntimeException(
              "Invalid map field number: " + number);
      }
    }
    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return burst.protocol.Burst.internal_static_protocol_BurstMessage_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              burst.protocol.BurstMessage.class, burst.protocol.BurstMessage.Builder.class);
    }

    // Construct using burst.protocol.BurstMessage.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      type_ = 0;

      internalGetMutableHeader().clear();
      data_ = com.google.protobuf.ByteString.EMPTY;

      serverPort_ = emptyIntList();
      bitField0_ = (bitField0_ & ~0x00000002);
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return burst.protocol.Burst.internal_static_protocol_BurstMessage_descriptor;
    }

    @java.lang.Override
    public burst.protocol.BurstMessage getDefaultInstanceForType() {
      return burst.protocol.BurstMessage.getDefaultInstance();
    }

    @java.lang.Override
    public burst.protocol.BurstMessage build() {
      burst.protocol.BurstMessage result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public burst.protocol.BurstMessage buildPartial() {
      burst.protocol.BurstMessage result = new burst.protocol.BurstMessage(this);
      int from_bitField0_ = bitField0_;
      result.type_ = type_;
      result.header_ = internalGetHeader();
      result.header_.makeImmutable();
      result.data_ = data_;
      if (((bitField0_ & 0x00000002) != 0)) {
        serverPort_.makeImmutable();
        bitField0_ = (bitField0_ & ~0x00000002);
      }
      result.serverPort_ = serverPort_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof burst.protocol.BurstMessage) {
        return mergeFrom((burst.protocol.BurstMessage)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(burst.protocol.BurstMessage other) {
      if (other == burst.protocol.BurstMessage.getDefaultInstance()) return this;
      if (other.type_ != 0) {
        setTypeValue(other.getTypeValue());
      }
      internalGetMutableHeader().mergeFrom(
          other.internalGetHeader());
      if (other.getData() != com.google.protobuf.ByteString.EMPTY) {
        setData(other.getData());
      }
      if (!other.serverPort_.isEmpty()) {
        if (serverPort_.isEmpty()) {
          serverPort_ = other.serverPort_;
          bitField0_ = (bitField0_ & ~0x00000002);
        } else {
          ensureServerPortIsMutable();
          serverPort_.addAll(other.serverPort_);
        }
        onChanged();
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      burst.protocol.BurstMessage parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (burst.protocol.BurstMessage) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private int type_ = 0;
    /**
     * <code>.protocol.BurstType type = 1;</code>
     * @return The enum numeric value on the wire for type.
     */
    @java.lang.Override public int getTypeValue() {
      return type_;
    }
    /**
     * <code>.protocol.BurstType type = 1;</code>
     * @param value The enum numeric value on the wire for type to set.
     * @return This builder for chaining.
     */
    public Builder setTypeValue(int value) {
      
      type_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>.protocol.BurstType type = 1;</code>
     * @return The type.
     */
    @java.lang.Override
    public burst.protocol.BurstType getType() {
      @SuppressWarnings("deprecation")
      burst.protocol.BurstType result = burst.protocol.BurstType.valueOf(type_);
      return result == null ? burst.protocol.BurstType.UNRECOGNIZED : result;
    }
    /**
     * <code>.protocol.BurstType type = 1;</code>
     * @param value The type to set.
     * @return This builder for chaining.
     */
    public Builder setType(burst.protocol.BurstType value) {
      if (value == null) {
        throw new NullPointerException();
      }
      
      type_ = value.getNumber();
      onChanged();
      return this;
    }
    /**
     * <code>.protocol.BurstType type = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearType() {
      
      type_ = 0;
      onChanged();
      return this;
    }

    private com.google.protobuf.MapField<
        java.lang.Integer, com.google.protobuf.Any> header_;
    private com.google.protobuf.MapField<java.lang.Integer, com.google.protobuf.Any>
    internalGetHeader() {
      if (header_ == null) {
        return com.google.protobuf.MapField.emptyMapField(
            HeaderDefaultEntryHolder.defaultEntry);
      }
      return header_;
    }
    private com.google.protobuf.MapField<java.lang.Integer, com.google.protobuf.Any>
    internalGetMutableHeader() {
      onChanged();;
      if (header_ == null) {
        header_ = com.google.protobuf.MapField.newMapField(
            HeaderDefaultEntryHolder.defaultEntry);
      }
      if (!header_.isMutable()) {
        header_ = header_.copy();
      }
      return header_;
    }

    public int getHeaderCount() {
      return internalGetHeader().getMap().size();
    }
    /**
     * <code>map&lt;int32, .google.protobuf.Any&gt; header = 2;</code>
     */

    @java.lang.Override
    public boolean containsHeader(
        int key) {
      
      return internalGetHeader().getMap().containsKey(key);
    }
    /**
     * Use {@link #getHeaderMap()} instead.
     */
    @java.lang.Override
    @java.lang.Deprecated
    public java.util.Map<java.lang.Integer, com.google.protobuf.Any> getHeader() {
      return getHeaderMap();
    }
    /**
     * <code>map&lt;int32, .google.protobuf.Any&gt; header = 2;</code>
     */
    @java.lang.Override

    public java.util.Map<java.lang.Integer, com.google.protobuf.Any> getHeaderMap() {
      return internalGetHeader().getMap();
    }
    /**
     * <code>map&lt;int32, .google.protobuf.Any&gt; header = 2;</code>
     */
    @java.lang.Override

    public com.google.protobuf.Any getHeaderOrDefault(
        int key,
        com.google.protobuf.Any defaultValue) {
      
      java.util.Map<java.lang.Integer, com.google.protobuf.Any> map =
          internalGetHeader().getMap();
      return map.containsKey(key) ? map.get(key) : defaultValue;
    }
    /**
     * <code>map&lt;int32, .google.protobuf.Any&gt; header = 2;</code>
     */
    @java.lang.Override

    public com.google.protobuf.Any getHeaderOrThrow(
        int key) {
      
      java.util.Map<java.lang.Integer, com.google.protobuf.Any> map =
          internalGetHeader().getMap();
      if (!map.containsKey(key)) {
        throw new java.lang.IllegalArgumentException();
      }
      return map.get(key);
    }

    public Builder clearHeader() {
      internalGetMutableHeader().getMutableMap()
          .clear();
      return this;
    }
    /**
     * <code>map&lt;int32, .google.protobuf.Any&gt; header = 2;</code>
     */

    public Builder removeHeader(
        int key) {
      
      internalGetMutableHeader().getMutableMap()
          .remove(key);
      return this;
    }
    /**
     * Use alternate mutation accessors instead.
     */
    @java.lang.Deprecated
    public java.util.Map<java.lang.Integer, com.google.protobuf.Any>
    getMutableHeader() {
      return internalGetMutableHeader().getMutableMap();
    }
    /**
     * <code>map&lt;int32, .google.protobuf.Any&gt; header = 2;</code>
     */
    public Builder putHeader(
        int key,
        com.google.protobuf.Any value) {
      
      if (value == null) {
  throw new NullPointerException("map value");
}

      internalGetMutableHeader().getMutableMap()
          .put(key, value);
      return this;
    }
    /**
     * <code>map&lt;int32, .google.protobuf.Any&gt; header = 2;</code>
     */

    public Builder putAllHeader(
        java.util.Map<java.lang.Integer, com.google.protobuf.Any> values) {
      internalGetMutableHeader().getMutableMap()
          .putAll(values);
      return this;
    }

    private com.google.protobuf.ByteString data_ = com.google.protobuf.ByteString.EMPTY;
    /**
     * <code>bytes data = 3;</code>
     * @return The data.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString getData() {
      return data_;
    }
    /**
     * <code>bytes data = 3;</code>
     * @param value The data to set.
     * @return This builder for chaining.
     */
    public Builder setData(com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      data_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>bytes data = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearData() {
      
      data_ = getDefaultInstance().getData();
      onChanged();
      return this;
    }

    private com.google.protobuf.Internal.IntList serverPort_ = emptyIntList();
    private void ensureServerPortIsMutable() {
      if (!((bitField0_ & 0x00000002) != 0)) {
        serverPort_ = mutableCopy(serverPort_);
        bitField0_ |= 0x00000002;
       }
    }
    /**
     * <pre>
     * 当type = REMOVE_PROXY_INFO 时有效
     * </pre>
     *
     * <code>repeated int32 serverPort = 4;</code>
     * @return A list containing the serverPort.
     */
    public java.util.List<java.lang.Integer>
        getServerPortList() {
      return ((bitField0_ & 0x00000002) != 0) ?
               java.util.Collections.unmodifiableList(serverPort_) : serverPort_;
    }
    /**
     * <pre>
     * 当type = REMOVE_PROXY_INFO 时有效
     * </pre>
     *
     * <code>repeated int32 serverPort = 4;</code>
     * @return The count of serverPort.
     */
    public int getServerPortCount() {
      return serverPort_.size();
    }
    /**
     * <pre>
     * 当type = REMOVE_PROXY_INFO 时有效
     * </pre>
     *
     * <code>repeated int32 serverPort = 4;</code>
     * @param index The index of the element to return.
     * @return The serverPort at the given index.
     */
    public int getServerPort(int index) {
      return serverPort_.getInt(index);
    }
    /**
     * <pre>
     * 当type = REMOVE_PROXY_INFO 时有效
     * </pre>
     *
     * <code>repeated int32 serverPort = 4;</code>
     * @param index The index to set the value at.
     * @param value The serverPort to set.
     * @return This builder for chaining.
     */
    public Builder setServerPort(
        int index, int value) {
      ensureServerPortIsMutable();
      serverPort_.setInt(index, value);
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 当type = REMOVE_PROXY_INFO 时有效
     * </pre>
     *
     * <code>repeated int32 serverPort = 4;</code>
     * @param value The serverPort to add.
     * @return This builder for chaining.
     */
    public Builder addServerPort(int value) {
      ensureServerPortIsMutable();
      serverPort_.addInt(value);
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 当type = REMOVE_PROXY_INFO 时有效
     * </pre>
     *
     * <code>repeated int32 serverPort = 4;</code>
     * @param values The serverPort to add.
     * @return This builder for chaining.
     */
    public Builder addAllServerPort(
        java.lang.Iterable<? extends java.lang.Integer> values) {
      ensureServerPortIsMutable();
      com.google.protobuf.AbstractMessageLite.Builder.addAll(
          values, serverPort_);
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 当type = REMOVE_PROXY_INFO 时有效
     * </pre>
     *
     * <code>repeated int32 serverPort = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearServerPort() {
      serverPort_ = emptyIntList();
      bitField0_ = (bitField0_ & ~0x00000002);
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:protocol.BurstMessage)
  }

  // @@protoc_insertion_point(class_scope:protocol.BurstMessage)
  private static final burst.protocol.BurstMessage DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new burst.protocol.BurstMessage();
  }

  public static burst.protocol.BurstMessage getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<BurstMessage>
      PARSER = new com.google.protobuf.AbstractParser<BurstMessage>() {
    @java.lang.Override
    public BurstMessage parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new BurstMessage(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<BurstMessage> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<BurstMessage> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public burst.protocol.BurstMessage getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

