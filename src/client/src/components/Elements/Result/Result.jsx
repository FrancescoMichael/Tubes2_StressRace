import React from "react";

export default function Result() {
    const data = [
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "2",
          "title": [
            'Michael',
            'Michael Jackson',
            'Michael Jordan',
            'Michael Jordan'
          ],
          "url": [
            'https://en.wikipedia.org/wiki/Michael',
            'https://en.wikipedia.org/wiki/Michael_Jackson',
            'https://en.wikipedia.org/wiki/Michael_Jordan',
            'https://en.wikipedia.org/wiki/Michael_Jordan'
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
      <div className= "d-flex justify-content-center mt-5 " style={{width: '90%'}}>
        {mappedData.length > 0 ? (
        <div>
          <div className="container-result text-white">
          {mappedData.map((path, i) => (
              <div key={i} className= "box-result flex flex-col justify-content-center mt-5 border-white border-2 p-3" style={{width: '30%'}}>
                  <h2 className="mx-auto" style={{ color: 'white', fontSize: '20px', display: 'inline-block'}}>PATH {i + 1}</h2>
                  <ul>
                      {path.map((link, j) => (
                      <div className= "flex flex-col justify-content-center mt-5" style={{width: '100%'}}>
                        <li key={j}>
                            <a className="mx-auto" style={{ color: 'white', fontSize: '20px', display: 'inline-block'}}href={link.value}>{link.label}</a>
                        </li>
                      </div>
                      ))}
                  </ul>
              </div>
          ))}
          </div>
        </div>
        ) : (
        <div>
          <h1 style={{ color: 'white', fontSize: '20px', display: 'inline-block'}}>NO RESULT</h1>
        </div>
        )}
      </div>
    );
}