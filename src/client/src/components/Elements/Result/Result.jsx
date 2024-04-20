import React from "react";

export default function Result() {
    const data = [
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express"
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express"
          ]
        },
        {
          "id": "2",
          "title": [
            'Michael',
            'Michael Jackson',
            'Michael Jordan',
            'Michael (archangel)',
            'Michelangelo',
            'Michael Schumacher',
            'Michael Sheen',
            'Michel Foucault',
            'Michael J. Fox',
            'Michael Phelps'
          ],
          "url": [
            'https://en.wikipedia.org/wiki/Michael',
            'https://en.wikipedia.org/wiki/Michael_Jackson',
            'https://en.wikipedia.org/wiki/Michael_Jordan',
            'https://en.wikipedia.org/wiki/Michael_(archangel)',
            'https://en.wikipedia.org/wiki/Michelangelo',
            'https://en.wikipedia.org/wiki/Michael_Schumacher',
            'https://en.wikipedia.org/wiki/Michael_Sheen',
            'https://en.wikipedia.org/wiki/Michel_Foucault',
            'https://en.wikipedia.org/wiki/Michael_J._Fox',
            'https://en.wikipedia.org/wiki/Michael_Phelps'
          ]
        }
    ];

    const mappedData = data.map(item => {
        const options = item.title.map((title, index) => ({
          label: title,
          value: item.url[index],
        }));
        return options;
      });
    return (
        <div className="text-white">
        {mappedData.map((path, i) => (
            <div key={i}>
                <h2>Path {i + 1}</h2>
                <ul>
                    {path.map((link, j) => (
                    <li key={j}>
                        <a href={link.value}>{link.label}</a>
                    </li>
                    ))}
                </ul>
            </div>
            
        ))}
        </div>
    );
}