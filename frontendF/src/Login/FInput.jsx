import "../App.css";
export const FInput = ({ setText }) => {
  const onChangeText = (e) => {
    setText(e.target.value);
  };
  return <input className="FInput" onChange={onChangeText} />;
};
export const FInput2 = ({ setText }) => {
  const onChangeText = (e) => {
    setText(e.target.value);
  };
  return <input className="FInput2" onChange={onChangeText} />;
};
export const TimeInput = ({ setTime }) => {
  const onChangeTime = (e) => {
    setTime(e.target.value);
    //console.log(new Date(e.target.value).toISOString());
  };
  return (
    <input
      className="FInput2"
      type={"datetime-local"}
      onChange={onChangeTime}
    />
  );
};
