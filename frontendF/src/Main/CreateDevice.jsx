import { useEffect, useState } from "react";
import Cookies from "universal-cookie";

import { FInput2 } from "../Login/FInput";
import { Loading } from "./Loading";

export const CreateDevice = ({ farm, setUFHandler, devices, deviceStatus }) => {
  const cookies = new Cookies();
  const [deviceSerial, setDeviceSerial] = useState("");
  const [devicePhoneNumber, setDevicePhoneNumber] = useState("");

  return (
    <div
      style={{
        backgroundColor: "#fff",
        width: "70%",
        borderRadius: 2,
        marginTop: 14,
        padding: 40
      }}
    >
      <h3
        style={{
          color: "#000000",
          borderBottomStyle: "solid",
          borderBottomColor: "#ccc",
          borderBottomWidth: 1,
          paddingBottom: 14,
          fontSize: "30px",
          marginTop: 0,
          marginBottom: 0
        }}
      >
        Devices
      </h3>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "flex-start",
        }}
      >
        <h4 style={{fontSize: 18}}>complete these information to create your device :</h4>
        <div
        style={{ display: "flex", flexDirection: "row", alignItems: "center" }}
      >
        <CInput setText={setDeviceSerial} title="device serial number : " />
        <CInput setText={setDevicePhoneNumber} title="device phone number : " />
        </div>
        <button
          className="chooseDeviceButton"
          onClick={() => {
            if ((deviceSerial.length > 0, devicePhoneNumber > 0)) {
              fetch(
                "http://asap.carriot.ir:8000/device/create",

                {
                  method: "POST",
                  headers: {
                    Authorization: "Bearer " + cookies.get("token"),
                    "content-type": "application/json",
                  },
                  body: JSON.stringify({
                    device_serial: deviceSerial,
                    phone: devicePhoneNumber,
                    farm_id: farm.id,
                  }),
                }
              ).then((res) => {
                if (res.status === 200) {
                  setUFHandler(res);
                  alert(
                    "device with serial number " +
                      deviceSerial +
                      " created successfully"
                  );
                } else {
                  alert("an error occurred\nERROR CODE : " + res.status);
                }
              });
            } else {
              alert("please fill the data");
            }
          }}
        >
          create device
        </button>
      </div>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          borderBottomStyle: "solid",
          borderBottomColor: "#ccc",
          borderBottomWidth: 1,
          fontSize: 18
        }}
      >
        {deviceStatus === 200 && devices.devices ? (
          devices.devices.length > 0 ? (
            <table width={"90%"} style={{ borderCollapse: "collapse", borderBlock: 1 }}>
              <tr
                style={{
                  borderBottomStyle: "dashed",
                  borderBottomWidth: 1,
                  fontWeight: "bold",
                }}
              >
                <td>serial</td>
                <td>phone number</td>
                <td>max hum.</td>
                <td>min hum</td>
              </tr>
              {devices.devices.map((item) => {
                const time = new Date(item.entry_time);
                return (
                  <tr>
                    <td>{item.device_serial}</td>
                    <td>{item.phone}</td>
                    <td>{item.max_humidity}</td>
                    <td>{item.min_humidity}</td>
                  </tr>
                );
              })}
            </table>
          ) : (
            <h4>
              no device found
              <h5>you can create a waterLog in the bottom of this page</h5>
            </h4>
          )
        ) : deviceStatus === 200 && !devices.devices ? (
          <h4>
            no device found
            <h5>you can create a waterLog in the bottom of this page</h5>
          </h4>
        ) : deviceStatus === 0 ? (
          <Loading />
        ) : (
          <h4>
            no device loaded
            <h5>an error occurred</h5>
          </h4>
        )}
      </div>
      
    </div>
  );
};
const CInput = ({ title, setText }) => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "row",
        alignItems: "center",
        fontSize: 16
      }}
    >
      <h5 style={{        fontSize: 16
}}>{title}</h5>
      <FInput2 setText={setText} />
    </div>
  );
};
