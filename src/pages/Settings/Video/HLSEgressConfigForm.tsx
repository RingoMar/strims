import React from "react";
import { useForm } from "react-hook-form";

import { HLSEgressConfig } from "../../../apis/strims/video/v1/hls_egress";
import { Button, ButtonSet, InputError, TextInput, ToggleInput } from "../../../components/Form";
import { useCall, useLazyCall } from "../../../contexts/FrontendApi";
import hostRegex from "../../../lib/hostRegex";

interface HLSEgressConfigFormData {
  enabled: boolean;
  serverAddr: string;
  publicServerAddr: string;
  serviceNetworkKeys: string[];
}

const HLSEgressConfigForm = () => {
  const [setConfigRes, setConfig] = useLazyCall("hlsEgress", "setConfig");

  const { handleSubmit, reset, control, formState } = useForm<HLSEgressConfigFormData>({
    mode: "onBlur",
    defaultValues: {
      enabled: false,
      serverAddr: "",
      publicServerAddr: "",
      serviceNetworkKeys: [],
    },
  });

  const setValues = ({ config }: { config?: HLSEgressConfig }) =>
    reset(
      {
        enabled: config.enabled,
        publicServerAddr: config.publicServerAddr,
      },
      {
        keepDirty: false,
        keepIsValid: false,
      }
    );

  useCall("hlsEgress", "getConfig", { onComplete: (res) => setValues(res) });

  const onSubmit = handleSubmit(async (data) => {
    const res = await setConfig({
      config: {
        enabled: data.enabled,
        publicServerAddr: data.publicServerAddr,
      },
    });
    setValues(res);
  });

  return (
    <>
      <form className="thing_form" onSubmit={onSubmit}>
        {setConfigRes.error && (
          <InputError error={setConfigRes.error.message || "Error saving ingress settings"} />
        )}
        <ToggleInput control={control} label="Enable" name="enabled" />
        <TextInput
          control={control}
          rules={{
            pattern: {
              value: hostRegex({
                localhost: false,
                strictPort: false,
              }),
              message: "Invalid address format",
            },
          }}
          label="Public address"
          description="Public address where clients can reach the HLS server."
          name="publicServerAddr"
          placeholder="ex: https://cdn.strims.gg"
        />
        <ButtonSet>
          <Button disabled={formState.isSubmitting || !formState.isDirty}>Save Changes</Button>
        </ButtonSet>
      </form>
    </>
  );
};

export default HLSEgressConfigForm;