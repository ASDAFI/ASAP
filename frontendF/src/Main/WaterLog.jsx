import { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import { Loading } from "./Loading";

import "./WaterLog.css";

export const WaterLog = () => {
  const cookies = new Cookies();
  const [waterLogs, setWaterLogs] = useState({});
  useEffect(() => {
    fetch("http://asap.carriot.ir:8000/water/logs", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + cookies.get("token"),
        "content-type": "application/json",
      },
    }).then((res) => {
      res.json().then((logs) => {
        setWaterLogs(logs);
      });
    });
  }, []);

  return (
    <div
      style={{
        backgroundColor: "#fff",
        width: "70%",
        borderRadius: 14,
        marginTop: 14,
        paddingLeft: 14,
      }}
    >
      <h3 style={{ color: "#00aaff" }}>Water Log</h3>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        {waterLogs.water_logs ? (
          waterLogs.water_logs.length > 0 ? (
            <table width={"90%"} style={{ borderCollapse: "collapse" }}>
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
                return (
                  <tr>
                    <td>{item.username}</td>
                    <td>{item.device_serial}</td>
                    <td>
                      {item.entry_time.substr(0, item.entry_time.indexOf("T")) +
                        " " +
                        item.entry_time.substr(
                          item.entry_time.indexOf("T") + 1,
                          8
                        )}
                    </td>

                    <td>{item.volume}</td>
                  </tr>
                );
              })}
            </table>
          ) : (
            <h4>
              no device found
              <h5>
                you can create a device by going to create device section from
                menu
              </h5>
            </h4>
          )
        ) : (
          <Loading />
        )}
      </div>
      {/* <div
        style={{ display: "flex", flexDirection: "row", alignItems: "center" }}
      >
        <h5>choose Device :</h5>
        <select className="waterLogSelect">
          {devices.devices.map((item) => {
            //console.log(item.device_serial);
            return <option>{item.device_serial}</option>;
          })}
        </select>
        <button
          className="chooseDeviceButton"
          onClick={}
        >
          choose
        </button> 
      </div>*/}
    </div>
  );
};
