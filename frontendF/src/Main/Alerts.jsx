import { useState } from "react";
import { useEffect } from "react";
import Cookies from "universal-cookie";
import { Loading } from "./Loading";
import "./WaterLog.css";

function GetData(cookies, setAlertStatus, setAlert) {
  fetch("http://asap.carriot.ir:8000/alerts/all", {
    method: "GET",
    headers: {
      Authorization: "Bearer " + cookies.get("token"),
      "content-type": "application/json",
    },
  }).then((res) => {
    if (res.status != 200) {
      alert("Error!");
      return;
    }
    setAlertStatus(res.status);
    res.json().then((alert) => {
      setAlert(alert);
    });
  });
}

export const Alerts = ({ }) => {
  const [alerts, setAlert] = useState({});
  const [AlertStatus, setAlertStatus] = useState(0);
  const cookies = new Cookies();
  useEffect(() => {
    GetData(cookies, setAlertStatus, setAlert);
    setInterval(() => {
      GetData(cookies, setAlertStatus, setAlert);
    }, 10000);
  }, []);

  return (
    <div
      style={{
        backgroundColor: "#fff",
        width: "70%",
        borderRadius: 2,
        marginTop: 14,
        padding: 40,
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
        }}
      >
        Alerts
      </h3>

      {/* <h5
        style={{
          borderBottomStyle: "solid",
          borderBottomColor: "#ccc",
          borderBottomWidth: 1,
          paddingBottom: 14,
          marginTop: 0,
        }}
      ></h5> */}
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          borderBottomStyle: "solid",
          borderBottomColor: "#ccc",
          borderBottomWidth: 1,
          fontSize: 18,
        }}
      >
        {/* {console.log(alerts)} */}
        {alerts.alerts ? (
          alerts.alerts.length > 0 ? (
            <table
              width={"90%"}
              style={{ borderCollapse: "collapse", borderBlock: 1 }}
            >
              <tr
                style={{
                  borderBottomStyle: "dashed",
                  borderBottomWidth: 1,
                  fontWeight: "bold",
                }}
              >
                <td>device serial</td>
                <td>Time</td>
                <td>text</td>
                <td>humidity</td>
              </tr>
              {alerts.alerts.map((item) => {
                const time = new Date(item.date);
                return (
                  <tr>
                    <td>{item.device_serial}</td>
                    <td>
                      {time.toLocaleDateString() +
                        " " +
                        time.toLocaleTimeString()}
                    </td>

                    <td>{item.text}</td>
                    <td>{item.humidity}</td>
                  </tr>
                );
              })}
            </table>
          ) : (
            <h4>no Alerts Found</h4>
          )
        ) : AlertStatus !== 200 ? (
          <h4>no Alerts found</h4>
        ) : AlertStatus === 0 ? (
          <Loading />
        ) : (
          <h4>
            no Alerts loaded
            <h5>an error occurred</h5>
          </h4>
        )}
      </div>
    </div>
  );
};
