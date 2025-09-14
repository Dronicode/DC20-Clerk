import { Link } from 'react-router-dom';

export default function Navbar() {
  return (
    <nav>
      <ul>
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/characters">Character Sheets</Link>
        </li>
        <li>
          <Link to="/">Other stuff coming later</Link>
        </li>
      </ul>
    </nav>
  );
}
