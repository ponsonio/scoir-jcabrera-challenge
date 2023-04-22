import React, { useState, version } from "react";

const LoginForm = (props) => {
  const [userInput, setUserInput] = useState({
    enteredUserName: "",
    enteredUserPassword: "",
  });

  const submitHandler = (event) => {
    event.preventDefault();
    props.onLoginHandler(userInput);
  };

  const userNameChangeHandler = (event) => {
    setUserInput((prev) => {
      return { ...prev, enteredUserName: event.target.value };
    });
  };

  const userPasswordChangeHandler = (event) => {
    setUserInput((prev) => {
      return { ...prev, enteredUserPassword: event.target.value };
    });
  };

  return (
    <form onSubmit={submitHandler}>
      <p> Input your credentials: </p>
      <div className="new-expense__controls">
        <div className="new-expense__control">
          <label>Username:</label>
          <input
            type="text"
            onChange={userNameChangeHandler}
            value={userInput.enteredUserName}
          />
        </div>
        <div className="new-expense__control">
          <label>Password</label>
          <input
            type="password"
            onChange={userPasswordChangeHandler}
            value={userInput.enteredUserPassword}
          />
        </div>
      </div>
      <div className="new-expense__actions">
        <button type="submit">Login</button>
      </div>
    </form>
  );
};

export default LoginForm;
