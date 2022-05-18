import Cookies from "universal-cookie";
import { useNavigate } from "react-router-dom";

import "./menu.css";
export const Menu = ({ currentPage, setCurrentPage, farm }) => {
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
        paddingRight: 14,
      }}
    >
      <div
        style={{
          borderBottomStyle: "solid",
          borderBottomColor: "#ccc",
          borderBottomWidth: 1,
        }}
      >
        <h4>{farm.name}</h4>
        <h6 style={{ marginTop: -14 }}>{farm.production} farm</h6>
      </div>
      <h5
        className="menuItem"
        style={{ color: currentPage === "waterLog" ? "rgb(0, 204, 0)" : "#111" }}
        onClick={() => {
          setCurrentPage("waterLog");
        }}
      >
        Water Log
      </h5>
      <h5
        className="menuItem"
        style={{ color: currentPage === "createDevice" ? "rgb(0, 204, 0)" : "#111" }}
        onClick={() => {
          setCurrentPage("createDevice");
        }}
      >
        Create Device
      </h5>
      <h5
        className="menuItem"
        style={{ color: currentPage === "report" ? "rgb(0, 204, 0)" : "#111" }}
        onClick={() => {
          setCurrentPage("report");
        }}
      >
        Report
      </h5>
      <h5
        className="menuItem"
        style={{ color: currentPage === "config" ? "rgb(0, 204, 0)" : "#111" }}
        onClick={() => {
          setCurrentPage("config");
        }}
      >
        Report
      </h5>
      <h5
        className="menuItem"
        style={{
          color: "#ff0040",
          borderStyle: "solid",
          borderWidth: 1,
          borderColor: "#ff0040",
          borderRadius: 9,
          padding: 6,
          display: "inline-block",
          marginTop: -4,
          marginLeft: 8
        }}
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
