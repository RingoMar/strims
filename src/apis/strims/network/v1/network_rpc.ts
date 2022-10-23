import strims_rpc_Host, { UnaryCallOptions as strims_rpc_UnaryCallOptions } from "@memelabs/protobuf/lib/rpc/host";
import strims_rpc_Service from "@memelabs/protobuf/lib/rpc/service";
import { Call as strims_rpc_Call } from "@memelabs/protobuf/lib/apis/strims/rpc/rpc";
import { Readable as GenericReadable } from "@memelabs/protobuf/lib/rpc/stream";

import {
  strims_network_v1_ICreateServerRequest,
  strims_network_v1_CreateServerRequest,
  strims_network_v1_CreateServerResponse,
  strims_network_v1_IUpdateServerConfigRequest,
  strims_network_v1_UpdateServerConfigRequest,
  strims_network_v1_UpdateServerConfigResponse,
  strims_network_v1_IDeleteNetworkRequest,
  strims_network_v1_DeleteNetworkRequest,
  strims_network_v1_DeleteNetworkResponse,
  strims_network_v1_IGetNetworkRequest,
  strims_network_v1_GetNetworkRequest,
  strims_network_v1_GetNetworkResponse,
  strims_network_v1_IListNetworksRequest,
  strims_network_v1_ListNetworksRequest,
  strims_network_v1_ListNetworksResponse,
  strims_network_v1_ICreateInvitationRequest,
  strims_network_v1_CreateInvitationRequest,
  strims_network_v1_CreateInvitationResponse,
  strims_network_v1_ICreateNetworkFromInvitationRequest,
  strims_network_v1_CreateNetworkFromInvitationRequest,
  strims_network_v1_CreateNetworkFromInvitationResponse,
  strims_network_v1_IWatchNetworksRequest,
  strims_network_v1_WatchNetworksRequest,
  strims_network_v1_WatchNetworksResponse,
  strims_network_v1_IUpdateDisplayOrderRequest,
  strims_network_v1_UpdateDisplayOrderRequest,
  strims_network_v1_UpdateDisplayOrderResponse,
  strims_network_v1_IUpdateAliasRequest,
  strims_network_v1_UpdateAliasRequest,
  strims_network_v1_UpdateAliasResponse,
  strims_network_v1_IGetUIConfigRequest,
  strims_network_v1_GetUIConfigRequest,
  strims_network_v1_GetUIConfigResponse,
  strims_network_v1_IListPeersRequest,
  strims_network_v1_ListPeersRequest,
  strims_network_v1_ListPeersResponse,
  strims_network_v1_IGrantPeerInvitationRequest,
  strims_network_v1_GrantPeerInvitationRequest,
  strims_network_v1_GrantPeerInvitationResponse,
  strims_network_v1_ITogglePeerBanRequest,
  strims_network_v1_TogglePeerBanRequest,
  strims_network_v1_TogglePeerBanResponse,
  strims_network_v1_IResetPeerRenameCooldownRequest,
  strims_network_v1_ResetPeerRenameCooldownRequest,
  strims_network_v1_ResetPeerRenameCooldownResponse,
  strims_network_v1_IDeletePeerRequest,
  strims_network_v1_DeletePeerRequest,
  strims_network_v1_DeletePeerResponse,
  strims_network_v1_IListAliasReservationsRequest,
  strims_network_v1_ListAliasReservationsRequest,
  strims_network_v1_ListAliasReservationsResponse,
  strims_network_v1_IResetAliasReservationCooldownRequest,
  strims_network_v1_ResetAliasReservationCooldownRequest,
  strims_network_v1_ResetAliasReservationCooldownResponse,
} from "./network";

export interface NetworkFrontendService {
  createServer(req: strims_network_v1_CreateServerRequest, call: strims_rpc_Call): Promise<strims_network_v1_CreateServerResponse> | strims_network_v1_CreateServerResponse;
  updateServerConfig(req: strims_network_v1_UpdateServerConfigRequest, call: strims_rpc_Call): Promise<strims_network_v1_UpdateServerConfigResponse> | strims_network_v1_UpdateServerConfigResponse;
  delete(req: strims_network_v1_DeleteNetworkRequest, call: strims_rpc_Call): Promise<strims_network_v1_DeleteNetworkResponse> | strims_network_v1_DeleteNetworkResponse;
  get(req: strims_network_v1_GetNetworkRequest, call: strims_rpc_Call): Promise<strims_network_v1_GetNetworkResponse> | strims_network_v1_GetNetworkResponse;
  list(req: strims_network_v1_ListNetworksRequest, call: strims_rpc_Call): Promise<strims_network_v1_ListNetworksResponse> | strims_network_v1_ListNetworksResponse;
  createInvitation(req: strims_network_v1_CreateInvitationRequest, call: strims_rpc_Call): Promise<strims_network_v1_CreateInvitationResponse> | strims_network_v1_CreateInvitationResponse;
  createNetworkFromInvitation(req: strims_network_v1_CreateNetworkFromInvitationRequest, call: strims_rpc_Call): Promise<strims_network_v1_CreateNetworkFromInvitationResponse> | strims_network_v1_CreateNetworkFromInvitationResponse;
  watch(req: strims_network_v1_WatchNetworksRequest, call: strims_rpc_Call): GenericReadable<strims_network_v1_WatchNetworksResponse>;
  updateDisplayOrder(req: strims_network_v1_UpdateDisplayOrderRequest, call: strims_rpc_Call): Promise<strims_network_v1_UpdateDisplayOrderResponse> | strims_network_v1_UpdateDisplayOrderResponse;
  updateAlias(req: strims_network_v1_UpdateAliasRequest, call: strims_rpc_Call): Promise<strims_network_v1_UpdateAliasResponse> | strims_network_v1_UpdateAliasResponse;
  getUIConfig(req: strims_network_v1_GetUIConfigRequest, call: strims_rpc_Call): Promise<strims_network_v1_GetUIConfigResponse> | strims_network_v1_GetUIConfigResponse;
  listPeers(req: strims_network_v1_ListPeersRequest, call: strims_rpc_Call): Promise<strims_network_v1_ListPeersResponse> | strims_network_v1_ListPeersResponse;
  grantPeerInvitation(req: strims_network_v1_GrantPeerInvitationRequest, call: strims_rpc_Call): Promise<strims_network_v1_GrantPeerInvitationResponse> | strims_network_v1_GrantPeerInvitationResponse;
  togglePeerBan(req: strims_network_v1_TogglePeerBanRequest, call: strims_rpc_Call): Promise<strims_network_v1_TogglePeerBanResponse> | strims_network_v1_TogglePeerBanResponse;
  resetPeerRenameCooldown(req: strims_network_v1_ResetPeerRenameCooldownRequest, call: strims_rpc_Call): Promise<strims_network_v1_ResetPeerRenameCooldownResponse> | strims_network_v1_ResetPeerRenameCooldownResponse;
  deletePeer(req: strims_network_v1_DeletePeerRequest, call: strims_rpc_Call): Promise<strims_network_v1_DeletePeerResponse> | strims_network_v1_DeletePeerResponse;
  listAliasReservations(req: strims_network_v1_ListAliasReservationsRequest, call: strims_rpc_Call): Promise<strims_network_v1_ListAliasReservationsResponse> | strims_network_v1_ListAliasReservationsResponse;
  resetAliasReservationCooldown(req: strims_network_v1_ResetAliasReservationCooldownRequest, call: strims_rpc_Call): Promise<strims_network_v1_ResetAliasReservationCooldownResponse> | strims_network_v1_ResetAliasReservationCooldownResponse;
}

export class UnimplementedNetworkFrontendService implements NetworkFrontendService {
  createServer(req: strims_network_v1_CreateServerRequest, call: strims_rpc_Call): Promise<strims_network_v1_CreateServerResponse> | strims_network_v1_CreateServerResponse { throw new Error("not implemented"); }
  updateServerConfig(req: strims_network_v1_UpdateServerConfigRequest, call: strims_rpc_Call): Promise<strims_network_v1_UpdateServerConfigResponse> | strims_network_v1_UpdateServerConfigResponse { throw new Error("not implemented"); }
  delete(req: strims_network_v1_DeleteNetworkRequest, call: strims_rpc_Call): Promise<strims_network_v1_DeleteNetworkResponse> | strims_network_v1_DeleteNetworkResponse { throw new Error("not implemented"); }
  get(req: strims_network_v1_GetNetworkRequest, call: strims_rpc_Call): Promise<strims_network_v1_GetNetworkResponse> | strims_network_v1_GetNetworkResponse { throw new Error("not implemented"); }
  list(req: strims_network_v1_ListNetworksRequest, call: strims_rpc_Call): Promise<strims_network_v1_ListNetworksResponse> | strims_network_v1_ListNetworksResponse { throw new Error("not implemented"); }
  createInvitation(req: strims_network_v1_CreateInvitationRequest, call: strims_rpc_Call): Promise<strims_network_v1_CreateInvitationResponse> | strims_network_v1_CreateInvitationResponse { throw new Error("not implemented"); }
  createNetworkFromInvitation(req: strims_network_v1_CreateNetworkFromInvitationRequest, call: strims_rpc_Call): Promise<strims_network_v1_CreateNetworkFromInvitationResponse> | strims_network_v1_CreateNetworkFromInvitationResponse { throw new Error("not implemented"); }
  watch(req: strims_network_v1_WatchNetworksRequest, call: strims_rpc_Call): GenericReadable<strims_network_v1_WatchNetworksResponse> { throw new Error("not implemented"); }
  updateDisplayOrder(req: strims_network_v1_UpdateDisplayOrderRequest, call: strims_rpc_Call): Promise<strims_network_v1_UpdateDisplayOrderResponse> | strims_network_v1_UpdateDisplayOrderResponse { throw new Error("not implemented"); }
  updateAlias(req: strims_network_v1_UpdateAliasRequest, call: strims_rpc_Call): Promise<strims_network_v1_UpdateAliasResponse> | strims_network_v1_UpdateAliasResponse { throw new Error("not implemented"); }
  getUIConfig(req: strims_network_v1_GetUIConfigRequest, call: strims_rpc_Call): Promise<strims_network_v1_GetUIConfigResponse> | strims_network_v1_GetUIConfigResponse { throw new Error("not implemented"); }
  listPeers(req: strims_network_v1_ListPeersRequest, call: strims_rpc_Call): Promise<strims_network_v1_ListPeersResponse> | strims_network_v1_ListPeersResponse { throw new Error("not implemented"); }
  grantPeerInvitation(req: strims_network_v1_GrantPeerInvitationRequest, call: strims_rpc_Call): Promise<strims_network_v1_GrantPeerInvitationResponse> | strims_network_v1_GrantPeerInvitationResponse { throw new Error("not implemented"); }
  togglePeerBan(req: strims_network_v1_TogglePeerBanRequest, call: strims_rpc_Call): Promise<strims_network_v1_TogglePeerBanResponse> | strims_network_v1_TogglePeerBanResponse { throw new Error("not implemented"); }
  resetPeerRenameCooldown(req: strims_network_v1_ResetPeerRenameCooldownRequest, call: strims_rpc_Call): Promise<strims_network_v1_ResetPeerRenameCooldownResponse> | strims_network_v1_ResetPeerRenameCooldownResponse { throw new Error("not implemented"); }
  deletePeer(req: strims_network_v1_DeletePeerRequest, call: strims_rpc_Call): Promise<strims_network_v1_DeletePeerResponse> | strims_network_v1_DeletePeerResponse { throw new Error("not implemented"); }
  listAliasReservations(req: strims_network_v1_ListAliasReservationsRequest, call: strims_rpc_Call): Promise<strims_network_v1_ListAliasReservationsResponse> | strims_network_v1_ListAliasReservationsResponse { throw new Error("not implemented"); }
  resetAliasReservationCooldown(req: strims_network_v1_ResetAliasReservationCooldownRequest, call: strims_rpc_Call): Promise<strims_network_v1_ResetAliasReservationCooldownResponse> | strims_network_v1_ResetAliasReservationCooldownResponse { throw new Error("not implemented"); }
}

export const registerNetworkFrontendService = (host: strims_rpc_Service, service: NetworkFrontendService): void => {
  host.registerMethod<strims_network_v1_CreateServerRequest, strims_network_v1_CreateServerResponse>("strims.network.v1.NetworkFrontend.CreateServer", service.createServer.bind(service), strims_network_v1_CreateServerRequest);
  host.registerMethod<strims_network_v1_UpdateServerConfigRequest, strims_network_v1_UpdateServerConfigResponse>("strims.network.v1.NetworkFrontend.UpdateServerConfig", service.updateServerConfig.bind(service), strims_network_v1_UpdateServerConfigRequest);
  host.registerMethod<strims_network_v1_DeleteNetworkRequest, strims_network_v1_DeleteNetworkResponse>("strims.network.v1.NetworkFrontend.Delete", service.delete.bind(service), strims_network_v1_DeleteNetworkRequest);
  host.registerMethod<strims_network_v1_GetNetworkRequest, strims_network_v1_GetNetworkResponse>("strims.network.v1.NetworkFrontend.Get", service.get.bind(service), strims_network_v1_GetNetworkRequest);
  host.registerMethod<strims_network_v1_ListNetworksRequest, strims_network_v1_ListNetworksResponse>("strims.network.v1.NetworkFrontend.List", service.list.bind(service), strims_network_v1_ListNetworksRequest);
  host.registerMethod<strims_network_v1_CreateInvitationRequest, strims_network_v1_CreateInvitationResponse>("strims.network.v1.NetworkFrontend.CreateInvitation", service.createInvitation.bind(service), strims_network_v1_CreateInvitationRequest);
  host.registerMethod<strims_network_v1_CreateNetworkFromInvitationRequest, strims_network_v1_CreateNetworkFromInvitationResponse>("strims.network.v1.NetworkFrontend.CreateNetworkFromInvitation", service.createNetworkFromInvitation.bind(service), strims_network_v1_CreateNetworkFromInvitationRequest);
  host.registerMethod<strims_network_v1_WatchNetworksRequest, strims_network_v1_WatchNetworksResponse>("strims.network.v1.NetworkFrontend.Watch", service.watch.bind(service), strims_network_v1_WatchNetworksRequest);
  host.registerMethod<strims_network_v1_UpdateDisplayOrderRequest, strims_network_v1_UpdateDisplayOrderResponse>("strims.network.v1.NetworkFrontend.UpdateDisplayOrder", service.updateDisplayOrder.bind(service), strims_network_v1_UpdateDisplayOrderRequest);
  host.registerMethod<strims_network_v1_UpdateAliasRequest, strims_network_v1_UpdateAliasResponse>("strims.network.v1.NetworkFrontend.UpdateAlias", service.updateAlias.bind(service), strims_network_v1_UpdateAliasRequest);
  host.registerMethod<strims_network_v1_GetUIConfigRequest, strims_network_v1_GetUIConfigResponse>("strims.network.v1.NetworkFrontend.GetUIConfig", service.getUIConfig.bind(service), strims_network_v1_GetUIConfigRequest);
  host.registerMethod<strims_network_v1_ListPeersRequest, strims_network_v1_ListPeersResponse>("strims.network.v1.NetworkFrontend.ListPeers", service.listPeers.bind(service), strims_network_v1_ListPeersRequest);
  host.registerMethod<strims_network_v1_GrantPeerInvitationRequest, strims_network_v1_GrantPeerInvitationResponse>("strims.network.v1.NetworkFrontend.GrantPeerInvitation", service.grantPeerInvitation.bind(service), strims_network_v1_GrantPeerInvitationRequest);
  host.registerMethod<strims_network_v1_TogglePeerBanRequest, strims_network_v1_TogglePeerBanResponse>("strims.network.v1.NetworkFrontend.TogglePeerBan", service.togglePeerBan.bind(service), strims_network_v1_TogglePeerBanRequest);
  host.registerMethod<strims_network_v1_ResetPeerRenameCooldownRequest, strims_network_v1_ResetPeerRenameCooldownResponse>("strims.network.v1.NetworkFrontend.ResetPeerRenameCooldown", service.resetPeerRenameCooldown.bind(service), strims_network_v1_ResetPeerRenameCooldownRequest);
  host.registerMethod<strims_network_v1_DeletePeerRequest, strims_network_v1_DeletePeerResponse>("strims.network.v1.NetworkFrontend.DeletePeer", service.deletePeer.bind(service), strims_network_v1_DeletePeerRequest);
  host.registerMethod<strims_network_v1_ListAliasReservationsRequest, strims_network_v1_ListAliasReservationsResponse>("strims.network.v1.NetworkFrontend.ListAliasReservations", service.listAliasReservations.bind(service), strims_network_v1_ListAliasReservationsRequest);
  host.registerMethod<strims_network_v1_ResetAliasReservationCooldownRequest, strims_network_v1_ResetAliasReservationCooldownResponse>("strims.network.v1.NetworkFrontend.ResetAliasReservationCooldown", service.resetAliasReservationCooldown.bind(service), strims_network_v1_ResetAliasReservationCooldownRequest);
}

export class NetworkFrontendClient {
  constructor(private readonly host: strims_rpc_Host) {}

  public createServer(req?: strims_network_v1_ICreateServerRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_CreateServerResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.CreateServer", new strims_network_v1_CreateServerRequest(req)), strims_network_v1_CreateServerResponse, opts);
  }

  public updateServerConfig(req?: strims_network_v1_IUpdateServerConfigRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_UpdateServerConfigResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.UpdateServerConfig", new strims_network_v1_UpdateServerConfigRequest(req)), strims_network_v1_UpdateServerConfigResponse, opts);
  }

  public delete(req?: strims_network_v1_IDeleteNetworkRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_DeleteNetworkResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.Delete", new strims_network_v1_DeleteNetworkRequest(req)), strims_network_v1_DeleteNetworkResponse, opts);
  }

  public get(req?: strims_network_v1_IGetNetworkRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_GetNetworkResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.Get", new strims_network_v1_GetNetworkRequest(req)), strims_network_v1_GetNetworkResponse, opts);
  }

  public list(req?: strims_network_v1_IListNetworksRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_ListNetworksResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.List", new strims_network_v1_ListNetworksRequest(req)), strims_network_v1_ListNetworksResponse, opts);
  }

  public createInvitation(req?: strims_network_v1_ICreateInvitationRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_CreateInvitationResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.CreateInvitation", new strims_network_v1_CreateInvitationRequest(req)), strims_network_v1_CreateInvitationResponse, opts);
  }

  public createNetworkFromInvitation(req?: strims_network_v1_ICreateNetworkFromInvitationRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_CreateNetworkFromInvitationResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.CreateNetworkFromInvitation", new strims_network_v1_CreateNetworkFromInvitationRequest(req)), strims_network_v1_CreateNetworkFromInvitationResponse, opts);
  }

  public watch(req?: strims_network_v1_IWatchNetworksRequest): GenericReadable<strims_network_v1_WatchNetworksResponse> {
    return this.host.expectMany(this.host.call("strims.network.v1.NetworkFrontend.Watch", new strims_network_v1_WatchNetworksRequest(req)), strims_network_v1_WatchNetworksResponse);
  }

  public updateDisplayOrder(req?: strims_network_v1_IUpdateDisplayOrderRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_UpdateDisplayOrderResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.UpdateDisplayOrder", new strims_network_v1_UpdateDisplayOrderRequest(req)), strims_network_v1_UpdateDisplayOrderResponse, opts);
  }

  public updateAlias(req?: strims_network_v1_IUpdateAliasRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_UpdateAliasResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.UpdateAlias", new strims_network_v1_UpdateAliasRequest(req)), strims_network_v1_UpdateAliasResponse, opts);
  }

  public getUIConfig(req?: strims_network_v1_IGetUIConfigRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_GetUIConfigResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.GetUIConfig", new strims_network_v1_GetUIConfigRequest(req)), strims_network_v1_GetUIConfigResponse, opts);
  }

  public listPeers(req?: strims_network_v1_IListPeersRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_ListPeersResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.ListPeers", new strims_network_v1_ListPeersRequest(req)), strims_network_v1_ListPeersResponse, opts);
  }

  public grantPeerInvitation(req?: strims_network_v1_IGrantPeerInvitationRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_GrantPeerInvitationResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.GrantPeerInvitation", new strims_network_v1_GrantPeerInvitationRequest(req)), strims_network_v1_GrantPeerInvitationResponse, opts);
  }

  public togglePeerBan(req?: strims_network_v1_ITogglePeerBanRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_TogglePeerBanResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.TogglePeerBan", new strims_network_v1_TogglePeerBanRequest(req)), strims_network_v1_TogglePeerBanResponse, opts);
  }

  public resetPeerRenameCooldown(req?: strims_network_v1_IResetPeerRenameCooldownRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_ResetPeerRenameCooldownResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.ResetPeerRenameCooldown", new strims_network_v1_ResetPeerRenameCooldownRequest(req)), strims_network_v1_ResetPeerRenameCooldownResponse, opts);
  }

  public deletePeer(req?: strims_network_v1_IDeletePeerRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_DeletePeerResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.DeletePeer", new strims_network_v1_DeletePeerRequest(req)), strims_network_v1_DeletePeerResponse, opts);
  }

  public listAliasReservations(req?: strims_network_v1_IListAliasReservationsRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_ListAliasReservationsResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.ListAliasReservations", new strims_network_v1_ListAliasReservationsRequest(req)), strims_network_v1_ListAliasReservationsResponse, opts);
  }

  public resetAliasReservationCooldown(req?: strims_network_v1_IResetAliasReservationCooldownRequest, opts?: strims_rpc_UnaryCallOptions): Promise<strims_network_v1_ResetAliasReservationCooldownResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.NetworkFrontend.ResetAliasReservationCooldown", new strims_network_v1_ResetAliasReservationCooldownRequest(req)), strims_network_v1_ResetAliasReservationCooldownResponse, opts);
  }
}

