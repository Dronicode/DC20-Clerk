import { createBrowserRouter } from 'react-router-dom';
import { AppLayout } from '@widgets/layout';
import { Home } from '@pages/home';
import { CharacterSheetManager } from '@pages/characterSheetManager';

export const router = createBrowserRouter([
  {
    path: '/',
    element: <AppLayout />,
    children: [
      { index: true, element: <Home /> },
      { path: 'characters', element: <CharacterSheetManager /> },
    ],
  },
]);
