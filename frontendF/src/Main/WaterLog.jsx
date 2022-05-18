import { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import { FInput2, TimeInput } from "../Login/FInput";
import { Loading } from "./Loading";

import "./WaterLog.css";

export const WaterLog = ({ devices }) => {
  const cookies = new Cookies();
  const [waterLogs, setWaterLogs] = useState({});
  const [volume, setVolume] = useState("");
  const [time, setTime] = useState("");
  const [ufCheck, setUFCheck] = useState({});
  const [logStatus, setLogStatus] = useState(0);

  useEffect(() => {
    fetch("http://asap.carriot.ir:8000/water/logs", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + cookies.get("token"),
        "content-type": "application/json",
      },
    }).then((res) => {
      setLogStatus(res.status);
      res.json().then((logs) => {
        setWaterLogs(logs);
        console.log("logs", logs);
      });
    });
  }, [ufCheck]);

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
          marginTop: 0
        }}
      >
        Water Log
      </h3>
      <h4 style={{ marginTop: 14, marginBottom: 2 , fontSize: "22px"}}>add waterLog</h4>
      <div
        style={{ display: "flex", flexDirection: "row", alignItems: "center" }}
      >
        <h5 style={{ fontSize: "16px" }}>Device</h5>
        <select id="waterLogDevice" className="waterLogSelect">
          {devices.devices ? (
            devices.devices.map((item) => {
              //console.log(item.device_serial);
              return (
                <option value={item.device_serial}>{item.device_serial}</option>
              );
            })
          ) : (
            <option value="">no device</option>
          )}
        </select>
        
        <h5 style={{ fontSize: "16px" }}>Volume</h5>
        <FInput2 setText={setVolume} />
        <h5 style={{ fontSize: "16px" }}>Time</h5>
        <TimeInput setTime={setTime} />
      
        
      </div>
      <button
          className="chooseDeviceButton"
          style={{marginBottom: 10}}
          onClick={() => {
            if (document.getElementById("waterLogDevice").value.length > 0) {
              if (volume.length > 0 && time.length > 0) {
                fetch(
                  "http://asap.carriot.ir:8000/water/logs/create/serial",

                  {
                    method: "POST",
                    headers: {
                      Authorization: "Bearer " + cookies.get("token"),
                      "content-type": "application/json",
                    },
                    body: JSON.stringify({
                      device_serial:
                        document.getElementById("waterLogDevice").value,
                      entry_time: new Date(time).toISOString(),
                      volume: volume,
                    }),
                  }
                ).then((res) => {
                  if (res.status === 200) {
                    alert("waterLog added successfully");
                    setUFCheck(res);
                  } else {
                    alert("an error occurred\nERROR CODE : " + res.status);
                  }
                  //console.log(res);
                });
              } else {
                alert("please fill the data");
              }
            } else {
              alert(
                "you have no devices\nyou can create devices from the menu"
              );
            }
          }}
        >
          add waterLog
        </button>
        <h5
        style={{
          borderBottomStyle: "solid",
          borderBottomColor: "#ccc",
          borderBottomWidth: 1,
          paddingBottom: 14,
          marginTop:0
        }}
      >
      </h5>
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
        {logStatus === 200 && waterLogs.water_logs ? (
          waterLogs.water_logs.length > 0 ? (
            <table width={"90%"} style={{ borderCollapse: "collapse", borderBlock: 1}}>
              <tr
                style={{
                  borderBottomStyle: "dashed",
                  borderBottomWidth: 1,
                  fontWeight: "bold",
                }}
              >
                <td>username</td>
                <td>device serial</td>
                <td>entry time</td>
                <td>volume</td>
              </tr>
              {waterLogs.water_logs.map((item) => {
                const time = new Date(item.entry_time);
                return (
                  <tr>
                    <td>{item.username}</td>
                    <td>{item.device_serial}</td>
                    <td>
                      {time.toLocaleDateString() +
                        " " +
                        time.toLocaleTimeString()}
                    </td>

                    <td>{item.volume}</td>
                  </tr>
                );
              })}
            </table>
          ) : (
            <h4>
              no waterLog found
              <h5>you can create a waterLog in the bottom of this page</h5>
            </h4>
          )
        ) : logStatus === 200 && !waterLogs.water_logs ? (
          <h4>
            no waterLog found
            <h5>you can create a waterLog in the bottom of this page</h5>
          </h4>
        ) : logStatus === 0 ? (
          <Loading />
        ) : (
          <h4>
            no waterLog loaded
            <h5>an error occurred</h5>
          </h4>
        )}
      </div>
      
    </div>
  );
};
