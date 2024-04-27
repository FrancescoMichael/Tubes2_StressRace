import { GraphCanvas } from "reagraph";

export default function Canvas ({ dataResult }) {
    // Initialize nodes and edges arrays
    let nodes = [];
    let edges = [];
    let existingNodes = {};

    // Iterate over each item in the data array
    dataResult.forEach((item, index) => {
        const titles = item.title;
        let prevNodeId = null;
    
        // Generate nodes and check for uniqueness
        titles.forEach((title, i) => {
            let nodeId = existingNodes[title];
            if (!nodeId) {
                nodeId = `n-${title}`;
                const nodeLabel = title;
                const node = { id: nodeId, label: nodeLabel };
                nodes.push(node);
                existingNodes[title] = nodeId;
            }
    
            // Add edges if previous node exists
            if (prevNodeId) {
                const source = prevNodeId;
                const target = nodeId;
                const label = `Edge ${source}->${target}`;
                const edgeId = `${source}->${target}`;
    
                // Create the edge object and push it to the edges array
                const edge = { id: edgeId, source, target, label };
                edges.push(edge);
            }
    
            prevNodeId = nodeId;
        });
    });
    return (
        <div className="graf graph-container border border-8 rounded-xl" style={{ border: "2px solid white" }}>
            <GraphCanvas nodes = {nodes} edges = {edges}/>
        </div>
    );
}