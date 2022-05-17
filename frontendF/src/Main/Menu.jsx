import Cookies from "universal-cookie";
import { useNavigate } from "react-router-dom";

import "./menu.css";
import { ReportLog } from "./chart";
export const Menu = () => {
  const cookies = new Cookies();
  const navigate = useNavigate();

  return (
    <div
      style={{
        backgroundColor: "#fff",
        width: "20%",
        borderRadius: 14,
        marginTop: 14,
        paddingLeft: 14,
      }}
    >
      <h5 className="menuItem" style={{ color: "#00aaff" }} onClick={event =>  window.location.href='/main/water'}>
        Water Log
      </h5>

      <h5 className="menuItem" style={{ color: "#00aaff" }} onClick={event =>  window.location.href='/main/report'}>
        Humidity Report
      </h5>
      <h5
        className="menuItem"
        style={{ color: "#111" }}
        onClick={() => {
          alert("under construction");
        }}
      >
        Create Device
      </h5>
      <h5
        className="menuItem"
        style={{ color: "#ff0040" }}
        onClick={() => {
          fetch("http://asap.carriot.ir:8000/logout", {
            method: "POST",
            headers: {
              Authorization: "Bearer " + cookies.get("token"),
              "content-type": "application/json",
            },
          }).then((res) => {
            if (res.status === 200) {
              cookies.remove("token");
              navigate("/");
            } else {
              alert("an error occurred");
            }
          });
        }}
      >
        Logout
      </h5>
    </div>
  );
};
