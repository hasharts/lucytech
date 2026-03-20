import React, { useState } from "react";

export default function App() {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);

  const analyze = async () => {
    const url = document.getElementById("url").value;

    try {
      setError(null);

      const res = await fetch("http://localhost:8080/api/analyze", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ url }),
      });

      if (!res.ok) {
        throw new Error(await res.text());
      }

      const json = await res.json();
      setData(json);
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div style={{ padding: "20px" }}>
      <h1>Web Analyzer</h1>

      <input
        id="url"
        placeholder="Enter URL (https://example.com)"
        style={{ width: "300px", marginRight: "10px" }}
      />

      <button onClick={analyze}>Analyze</button>

      {error && <p style={{ color: "red" }}>{error}</p>}

      {data && (
        <div style={{ marginTop: "20px" }}>
          <h2>{data.title}</h2>
          <p>HTML Version: {data.htmlVersion}</p>
          <p>Internal Links: {data.internalLinks}</p>
          <p>External Links: {data.externalLinks}</p>
          <p>Login Form: {data.hasLogin ? "Yes" : "No"}</p>
        </div>
      )}
    </div>
  );
}