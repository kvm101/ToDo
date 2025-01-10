import streamlit as st
import numpy as np

query_row = np.random.randn(30, 1)
user_row = np.random.randn(10, 1)

st.title("Telemetry")
st.subheader("Query Payload")

queries = st.line_chart(query_row)

st.subheader("Users online")
users = st.line_chart(user_row)

def page1():
    st.write(st.session_state.foo)

def page2():
    st.write(st.session_state.bar)

pages = ([
        st.Page("users.py", title="Users"),
        st.Page("tasks.py", title="Tasks"),
        st.Page("productivity.py", title="Productivity"),
        st.Page("errors.py", title="Errors"),
])

pg = st.navigation(pages)