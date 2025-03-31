import React, { useState, useEffect } from "react";

const Tabs = ({ tabs }) => {
  const [selectedTab, setSelectedTab] = useState(null);

  const [componentToRender, setComponentToRender] = useState(null);

  // Set initial tab when component mounts
  useEffect(() => {
    const initialTab = tabs.find((tab) => tab.current);
    if (initialTab) {
      setSelectedTab(initialTab.name);
      setComponentToRender(initialTab.componentToRender);
    } else if (tabs.length > 0) {
      setSelectedTab(tabs[0].name);
      setComponentToRender(tabs[0].componentToRender);
    }
  }, [tabs]);

  const handleTabChange = (tabName) => {
    const selectedTabData = tabs.find((tab) => tab.name === tabName);

    setSelectedTab(selectedTabData.name);
    setComponentToRender(selectedTabData.componentToRender);
  };

  return (
    <div>
      {/* Render tabs */}
      <div className="hidden sm:block">
        <div className="border-b border-gray-200">
          <nav className="-mb-px flex space-x-8" aria-label="Tabs">
            {tabs.map((tab) => (
              <div
                key={tab.name}
                onClick={() => handleTabChange(tab.name)}
                className={`${
                  tab.name === selectedTab
                    ? "border-blue-500 text-blue-600"
                    : "border-transparent cursor-pointer text-gray-500 hover:border-gray-300 hover:text-gray-700"
                } group inline-flex items-center border-b-2 py-4 px-1 text-sm font-medium cursor-pointer`}
                aria-current={tab.name === selectedTab ? "page" : undefined}
              >
                <tab.icon
                  className={`${
                    tab.name === selectedTab
                      ? "text-blue-500"
                      : "text-gray-400 group-hover:text-gray-500"
                  } -ml-0.5 mr-2 h-5 w-5`}
                  aria-hidden="true"
                />
                <span>{tab.name}</span>
              </div>
            ))}
          </nav>
        </div>
      </div>

      <div className="tab-content py-4">{componentToRender}</div>
    </div>
  );
};

export default Tabs;
