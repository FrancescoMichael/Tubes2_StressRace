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
  // nodes: {
  //   shape: 'dot',
  //   scaling: {
  //     min: 10,
  //     max: 10,
  //     label: {
  //       min: 8,
  //       max: 30,
  //       drawThreshold: 12,
  //       maxVisible: 20
  //     }
  //   },
  //   color: {
  //     background: '#6D6D6D',
  //   },
  //   font: {
  //     color: '#FFFFFF',
  //     size: 12
  //   }
  // }
};

export default function GraphView({ dataResult }) {
  const nodes = [];
  const edges = [];

  if(dataResult.length > 12) {
    dataResult = dataResult.slice(0, 12);
  }
  
  dataResult.forEach(step => {
      if (step.title !== null) {
          step.title.forEach((title, index) => {
              if (!nodes.find(node => node.id === title)) {
                  nodes.push({ id: title, label: title, url: step.url ? step.url[index] : null });
              }
              if (index > 0 && step.title[index - 1] !== null) {
                  edges.push({ from: step.title[index - 1], to: title });
              }
          });
      }
  });

  const handleNodeSelect = ({ nodes }) => {
    console.log("Selected nodes:", nodes);
    if (nodes.length === 1) {
      const selectedNode = nodes[0];
      const nodeData = nodes.find(node => node.id === selectedNode);
      if (nodeData && nodeData.url) {
        window.open(nodeData.url, "_blank");
      }
    }
  };

  const [state, setState] = useState({
    graph: {
      nodes,
      edges
    },
    events: {
      select: handleNodeSelect
    }
  });

  const { graph, events } = state;

  return (
    <div className="graf graph-container border border-8 rounded-xl" style={{ border: "2px solid white" }}>
      <Graph graph={graph} options={options} events={events} style={{ height: "480px"}} />
    </div>
  );
}
