import i18n from "@/i18n";
import store from "@/store";
import { Outlet, useLocation } from "@ice/runtime"
import { useEffect } from "react";
import { CollectProviders } from "@knockout-js/layout";

export default () => {
  const [appState] = store.useModel('app'),
    location = useLocation();

  useEffect(() => {
    i18n.changeLanguage(appState.locale);
  }, [appState.locale]);

  return <CollectProviders
    locale={appState.locale}
    dark={appState.darkMode}
    pathname={location.pathname}
  >
    <Outlet />
  </CollectProviders>;
}
