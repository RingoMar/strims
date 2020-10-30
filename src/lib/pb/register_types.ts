// GENERATED CODE -- DO NOT EDIT!
import * as pb from "./pb";
import { registerType } from "./registry";

registerType("MonitorSwarmsRequest", pb.MonitorSwarmsRequest);
registerType("MonitorSwarmsResponse", pb.MonitorSwarmsResponse);
registerType("BootstrapDHTRequest", pb.BootstrapDHTRequest);
registerType("BootstrapDHTResponse", pb.BootstrapDHTResponse);
registerType("NegotiateWRTCRequest", pb.NegotiateWRTCRequest);
registerType("NegotiateWRTCResponse", pb.NegotiateWRTCResponse);
registerType("CreateChatServerRequest", pb.CreateChatServerRequest);
registerType("CreateChatServerResponse", pb.CreateChatServerResponse);
registerType("UpdateChatServerRequest", pb.UpdateChatServerRequest);
registerType("UpdateChatServerResponse", pb.UpdateChatServerResponse);
registerType("DeleteChatServerRequest", pb.DeleteChatServerRequest);
registerType("DeleteChatServerResponse", pb.DeleteChatServerResponse);
registerType("GetChatServerRequest", pb.GetChatServerRequest);
registerType("GetChatServerResponse", pb.GetChatServerResponse);
registerType("ListChatServersRequest", pb.ListChatServersRequest);
registerType("ListChatServersResponse", pb.ListChatServersResponse);
registerType("OpenChatServerRequest", pb.OpenChatServerRequest);
registerType("ChatServerEvent", pb.ChatServerEvent);
registerType("CallChatServerRequest", pb.CallChatServerRequest);
registerType("OpenChatClientRequest", pb.OpenChatClientRequest);
registerType("ChatClientEvent", pb.ChatClientEvent);
registerType("ChatRoom", pb.ChatRoom);
registerType("ChatServer", pb.ChatServer);
registerType("MessageEntities", pb.MessageEntities);
registerType("CallChatClientRequest", pb.CallChatClientRequest);
registerType("CallChatClientResponse", pb.CallChatClientResponse);
registerType("PProfRequest", pb.PProfRequest);
registerType("PProfResponse", pb.PProfResponse);
registerType("ReadMetricsRequest", pb.ReadMetricsRequest);
registerType("ReadMetricsResponse", pb.ReadMetricsResponse);
registerType("GetDirectoryEventsRequest", pb.GetDirectoryEventsRequest);
registerType("TestDirectoryPublishRequest", pb.TestDirectoryPublishRequest);
registerType("TestDirectoryPublishResponse", pb.TestDirectoryPublishResponse);
registerType("DirectoryListing", pb.DirectoryListing);
registerType("DirectoryServerEvent", pb.DirectoryServerEvent);
registerType("CallDirectoryServerRequest", pb.CallDirectoryServerRequest);
registerType("OpenDirectoryClientRequest", pb.OpenDirectoryClientRequest);
registerType("DirectoryClientEvent", pb.DirectoryClientEvent);
registerType("HashTableMessage", pb.HashTableMessage);
registerType("ServerConfig", pb.ServerConfig);
registerType("NickservNick", pb.NickservNick);
registerType("NickServToken", pb.NickServToken);
registerType("NickServRPCCommand", pb.NickServRPCCommand);
registerType("NickServRPCResponse", pb.NickServRPCResponse);
registerType("PeerIndexMessage", pb.PeerIndexMessage);
registerType("CreateProfileRequest", pb.CreateProfileRequest);
registerType("CreateProfileResponse", pb.CreateProfileResponse);
registerType("UpdateProfileRequest", pb.UpdateProfileRequest);
registerType("UpdateProfileResponse", pb.UpdateProfileResponse);
registerType("DeleteProfileRequest", pb.DeleteProfileRequest);
registerType("DeleteProfileResponse", pb.DeleteProfileResponse);
registerType("LoadProfileRequest", pb.LoadProfileRequest);
registerType("LoadProfileResponse", pb.LoadProfileResponse);
registerType("GetProfileRequest", pb.GetProfileRequest);
registerType("GetProfileResponse", pb.GetProfileResponse);
registerType("ListProfilesRequest", pb.ListProfilesRequest);
registerType("ListProfilesResponse", pb.ListProfilesResponse);
registerType("LoadSessionRequest", pb.LoadSessionRequest);
registerType("LoadSessionResponse", pb.LoadSessionResponse);
registerType("NetworkIcon", pb.NetworkIcon);
registerType("CreateNetworkRequest", pb.CreateNetworkRequest);
registerType("CreateNetworkResponse", pb.CreateNetworkResponse);
registerType("UpdateNetworkRequest", pb.UpdateNetworkRequest);
registerType("UpdateNetworkResponse", pb.UpdateNetworkResponse);
registerType("DeleteNetworkRequest", pb.DeleteNetworkRequest);
registerType("DeleteNetworkResponse", pb.DeleteNetworkResponse);
registerType("GetNetworkRequest", pb.GetNetworkRequest);
registerType("GetNetworkResponse", pb.GetNetworkResponse);
registerType("ListNetworksRequest", pb.ListNetworksRequest);
registerType("ListNetworksResponse", pb.ListNetworksResponse);
registerType("StorageKey", pb.StorageKey);
registerType("Key", pb.Key);
registerType("CertificateRequest", pb.CertificateRequest);
registerType("Certificate", pb.Certificate);
registerType("Profile", pb.Profile);
registerType("ProfileSummary", pb.ProfileSummary);
registerType("Network", pb.Network);
registerType("NetworkMembership", pb.NetworkMembership);
registerType("CreateNetworkInvitationRequest", pb.CreateNetworkInvitationRequest);
registerType("CreateNetworkInvitationResponse", pb.CreateNetworkInvitationResponse);
registerType("Invitation", pb.Invitation);
registerType("InvitationV0", pb.InvitationV0);
registerType("CreateNetworkFromInvitationRequest", pb.CreateNetworkFromInvitationRequest);
registerType("CreateNetworkFromInvitationResponse", pb.CreateNetworkFromInvitationResponse);
registerType("Mutex", pb.Mutex);
registerType("PubSubEvent", pb.PubSubEvent);
registerType("Call", pb.Call);
registerType("Error", pb.Error);
registerType("Cancel", pb.Cancel);
registerType("Undefined", pb.Undefined);
registerType("Close", pb.Close);
registerType("JoinSwarmRequest", pb.JoinSwarmRequest);
registerType("JoinSwarmResponse", pb.JoinSwarmResponse);
registerType("LeaveSwarmRequest", pb.LeaveSwarmRequest);
registerType("LeaveSwarmResponse", pb.LeaveSwarmResponse);
registerType("GetIngressStreamsRequest", pb.GetIngressStreamsRequest);
registerType("GetIngressStreamsResponse", pb.GetIngressStreamsResponse);
registerType("StartRTMPIngressRequest", pb.StartRTMPIngressRequest);
registerType("StartRTMPIngressResponse", pb.StartRTMPIngressResponse);
registerType("StartHLSEgressRequest", pb.StartHLSEgressRequest);
registerType("StartHLSEgressResponse", pb.StartHLSEgressResponse);
registerType("StopHLSEgressRequest", pb.StopHLSEgressRequest);
registerType("StopHLSEgressResponse", pb.StopHLSEgressResponse);
registerType("StartSwarmRequest", pb.StartSwarmRequest);
registerType("StartSwarmResponse", pb.StartSwarmResponse);
registerType("WriteToSwarmRequest", pb.WriteToSwarmRequest);
registerType("WriteToSwarmResponse", pb.WriteToSwarmResponse);
registerType("StopSwarmRequest", pb.StopSwarmRequest);
registerType("StopSwarmResponse", pb.StopSwarmResponse);
registerType("PublishSwarmRequest", pb.PublishSwarmRequest);
registerType("PublishSwarmResponse", pb.PublishSwarmResponse);
registerType("OpenVideoServerRequest", pb.OpenVideoServerRequest);
registerType("VideoServerOpenResponse", pb.VideoServerOpenResponse);
registerType("WriteToVideoServerRequest", pb.WriteToVideoServerRequest);
registerType("WriteToVideoServerResponse", pb.WriteToVideoServerResponse);
registerType("OpenVideoClientRequest", pb.OpenVideoClientRequest);
registerType("VideoClientEvent", pb.VideoClientEvent);
registerType("VideoClientCallRequest", pb.VideoClientCallRequest);
registerType("SwarmThingMessage", pb.SwarmThingMessage);
registerType("StartVPNRequest", pb.StartVPNRequest);
registerType("StartVPNResponse", pb.StartVPNResponse);
registerType("StopVPNRequest", pb.StopVPNRequest);
registerType("StopVPNResponse", pb.StopVPNResponse);
registerType("NetworkEvent", pb.NetworkEvent);
registerType("NetworkAddress", pb.NetworkAddress);
registerType("PeerInit", pb.PeerInit);
registerType("NetworkHandshake", pb.NetworkHandshake);
registerType("BrokerProxyRequest", pb.BrokerProxyRequest);
registerType("BrokerProxyEvent", pb.BrokerProxyEvent);
registerType("BrokerProxySendKeysRequest", pb.BrokerProxySendKeysRequest);
registerType("BrokerProxySendKeysResponse", pb.BrokerProxySendKeysResponse);
registerType("BrokerProxyReceiveKeysRequest", pb.BrokerProxyReceiveKeysRequest);
registerType("BrokerProxyReceiveKeysResponse", pb.BrokerProxyReceiveKeysResponse);
registerType("BrokerProxyDataRequest", pb.BrokerProxyDataRequest);
registerType("BrokerProxyDataResponse", pb.BrokerProxyDataResponse);
registerType("BrokerProxyCloseRequest", pb.BrokerProxyCloseRequest);
registerType("BrokerProxyCloseResponse", pb.BrokerProxyCloseResponse);
registerType("BootstrapClient", pb.BootstrapClient);
registerType("BootstrapClientWebSocketOptions", pb.BootstrapClientWebSocketOptions);
registerType("CreateBootstrapClientRequest", pb.CreateBootstrapClientRequest);
registerType("CreateBootstrapClientResponse", pb.CreateBootstrapClientResponse);
registerType("UpdateBootstrapClientRequest", pb.UpdateBootstrapClientRequest);
registerType("UpdateBootstrapClientResponse", pb.UpdateBootstrapClientResponse);
registerType("DeleteBootstrapClientRequest", pb.DeleteBootstrapClientRequest);
registerType("DeleteBootstrapClientResponse", pb.DeleteBootstrapClientResponse);
registerType("GetBootstrapClientRequest", pb.GetBootstrapClientRequest);
registerType("GetBootstrapClientResponse", pb.GetBootstrapClientResponse);
registerType("ListBootstrapClientsRequest", pb.ListBootstrapClientsRequest);
registerType("ListBootstrapClientsResponse", pb.ListBootstrapClientsResponse);
registerType("ListBootstrapPeersRequest", pb.ListBootstrapPeersRequest);
registerType("ListBootstrapPeersResponse", pb.ListBootstrapPeersResponse);
registerType("BootstrapPeer", pb.BootstrapPeer);
registerType("BootstrapServiceMessage", pb.BootstrapServiceMessage);
registerType("PublishNetworkToBootstrapPeerRequest", pb.PublishNetworkToBootstrapPeerRequest);
registerType("PublishNetworkToBootstrapPeerResponse", pb.PublishNetworkToBootstrapPeerResponse);
registerType("PeerExchangeMessage", pb.PeerExchangeMessage);
registerType("CAMessage", pb.CAMessage);
