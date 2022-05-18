import { useNavigate } from "react-router-dom";

import "./App.css";

const Intro = () => {
  const navigate = useNavigate();
  return (
    <div>
      <img
        src={require("./Assets/loginBG.jpg")}
        alt="bg"
        style={{
          width: "100%",
          position: "fixed",
          left: 0,
          top: 0,
          overflow: "hidden",
          zIndex: -1,
        }}
      />
      <div
        style={{
          backgroundColor: "#ffffff70",
          borderRadius: 80,
          width: "70vw",
          height: "50vh",
          marginTop: "8%",
          overflow: "hidden",
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <h1
          style={{
            fontSize: 40,
            color: "#004f12",
            fontFamily: "Lato",
            borderBottomWidth: 2,
            borderBottomStyle: "solid",
            marginBottom: 44,
          }}
        >
          Wellcome to the Farm Manager !
        </h1>
        <h3>Manage your farm !</h3>
        <button className="FButton" onClick={() => navigate("./login")}>
          Start Now!
        </button>
      </div>
    </div>
  );
};
export default Intro;
