"use client"
import React, { useState } from 'react';

const MyComponent = () => {
  const [isVisible, setIsVisible] = useState(false);

  const toggleVisibility = () => {
    setIsVisible(!isVisible);
  };

  return (
    <div>
      <button onClick={toggleVisibility} className="bg-blue-500 text-white font-bold py-2 px-4 rounded">
        {isVisible ? 'Hide' : 'Show'} Div
      </button>

      {/* Use conditional rendering to show or hide the div */}
      {isVisible && (
        <div className="bg-gray-200 p-4 mt-4 rounded">
          This is the content of the div.
        </div>
      )}
    </div>
  );
};

export default MyComponent;