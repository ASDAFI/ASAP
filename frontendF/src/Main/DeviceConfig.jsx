import { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import { FInput2, TimeInput } from "../Login/FInput";
import { Loading } from "./Loading";

import "./WaterLog.css";

export const DeviceConfig = ({ devices }) => {
    const cookies = new Cookies();
    const [minHumidity, setMinHumidity] = useState("");
    const [maxHumidity, setMaxHumidity] = useState("");
    const [time, setTime] = useState("");
    const [deviceSerial, setDeviceSerial] = useState("");
    const [ufCheck, setUFCheck] = useState({});
    const [logStatus, setLogStatus] = useState(0);
    const [waterLogs, setWaterLogs] = useState({});


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
                Device Config
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

                <h5 style={{ fontSize: "16px" }}>Minimum Humidity</h5>
                <FInput2 setText={setMinHumidity} />
                <h5 style={{ fontSize: "16px" }}>Maximum Humidity</h5>
                <FInput2 setText={setMaxHumidity} />

            </div>
            <button
                className="chooseDeviceButton"
                style={{marginBottom: 10}}
                onClick={() => {
                    if (document.getElementById("waterLogDevice").value.length > 0) {
                        if (minHumidity.length > 0 && maxHumidity.length > 0 && minHumidity<maxHumidity) {
                            fetch(
                                "http://asap.carriot.ir:8000/device/set",

                                {
                                    method: "POST",
                                    headers: {
                                        Authorization: "Bearer " + cookies.get("token"),
                                        "content-type": "application/json",
                                    },
                                    body: JSON.stringify({
                                        device_serial:
                                        document.getElementById("waterLogDevice").value,
                                        max_humidity: maxHumidity,
                                        min_humidity: minHumidity,
                                        date: new Date(time).toISOString(),
                                    }),
                                }
                            ).then((res) => {
                                if (res.status === 200) {
                                    alert("device was configured successfully");
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
                            "you have no devices\n"
                        );
                    }
                }}
            >
                configure
            </button>

        </div>
    );
};
