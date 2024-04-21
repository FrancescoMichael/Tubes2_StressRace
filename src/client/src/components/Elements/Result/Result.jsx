import React, { useState, useEffect } from "react";
import ResultList from "./ResultList";
import Pagination from "./Pagination";
import { getData } from "./DataResult.service"; // Import the getData function

export default function Result() {
  const [data, setData] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [postPerPage] = useState(6);

  useEffect(() => {
    getData((responseData) => {
      setData(responseData);
    });
  }, []);
  
  useEffect(() => {
    console.log("Ini data : ", data);
  }, [data]);
  

  const lastPostIndex = currentPage * postPerPage;
  const firstPostIndex = lastPostIndex - postPerPage;
  const currentPost = data.slice(firstPostIndex, lastPostIndex);

  return (
    <>
      <ResultList dataResults={currentPost} />
      <Pagination
        totalPosts={data.length}
        postPerPage={postPerPage}
        setCurrentPage={setCurrentPage}
        currentPage={currentPage}
      />
    </>
  );
}
