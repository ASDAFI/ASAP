export const Loading = () => {
  return (
    <div
      style={{
        width: 70,
        height: 70,
        borderRadius: 50,
        borderStyle: "dashed",
        overflow: "hidden",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        borderColor: "#009105",
        marginTop: "8%",
        marginBottom: "8%",
      }}
    >
      <img
        src={require("../Assets/97952-loading-animation-blue.gif")}
        alt=""
        style={{
          width: 100,
          height: 100,
        }}
      />
    </div>
  );
};
