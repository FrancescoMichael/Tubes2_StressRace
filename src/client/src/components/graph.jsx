import { width } from '@fortawesome/free-solid-svg-icons/fa0';
import React, { useState } from 'react';
import Graph from 'react-vis-network-graph';

const options = {
  layout: {
    hierarchical: false
  },
  edges: {
    color: "#FFFFFF"
  },
  nodes: {
    shape: 'dot',
    scaling: {
      min: 10,
      max: 10,
      label: {
        min: 8,
        max: 30,
        drawThreshold: 12,
        maxVisible: 20
      }
    },
    color: {
      background: '#6D6D6D',
    },
    font: {
      color: '#FFFFFF',
      size: 12
    }
  }
};
  
function randomColor() {
  const red = Math.floor(Math.random() * 256).toString(16).padStart(2, '0');
  const green = Math.floor(Math.random() * 256).toString(16).padStart(2, '0');
  const blue = Math.floor(Math.random() * 256).toString(16).padStart(2, '0');
  return `#${red}${green}${blue}`;
}

export default function GraphView({ dataResult }) {
  const nodes = [];
  const edges = [];
  
  // Populate nodes and edges from dataResult

  dataResult.forEach(step => {
    step.title.forEach((title, index) => {
      // Add node if it doesn't exist already
      if (!nodes.find(node => node.id === title)) {
        nodes.push({ id: title, label: title, url: step.url[index] });
      }

      // Add edge if not the first step
      if (index > 0) {
        edges.push({ from: step.title[index - 1], to: title });
      }
    });
  });
  
  // Define event handler for node selection
  const handleNodeSelect = ({ nodes }) => {
    console.log("Selected nodes:", nodes);
    // Open URL if a single node is selected
    if (nodes.length === 1) {
      const selectedNode = nodes[0];
      const nodeData = nodes.find(node => node.id === selectedNode);
      if (nodeData && nodeData.url) {
        window.open(nodeData.url, "_blank");
      }
    }
  };

  // Define state with graph and events
  const [state, setState] = useState({
    graph: {
      nodes,
      edges
    },
    events: {
      select: handleNodeSelect
    }
  });

  // Destructure state for graph and events
  const { graph, events } = state;

  return (
    <div className="graf graph-container border border-8 rounded-xl" style={{ border: "2px solid white" }}>
      <Graph graph={graph} options={options} events={events} style={{ height: "480px"}} />
    </div>
  );
}
